package convert

import (
	"encoding/json"
	"fmt"
	"strings"
)

func PostJSONToMarkdown(data []byte) (string, error) {
	doc, err := unmarshalJSON(data)
	if err != nil {
		return "", err
	}
	return contentToMarkdown(doc.Content), nil
}

type document struct {
	Type    string    `json:"type"`
	Content []content `json:"content"`
}

type content struct {
	Type    string    `json:"type"`
	Attrs   any       `json:"attrs,omitempty"`
	Content []content `json:"content,omitempty"`
	Text    string    `json:"text,omitempty"`
	Marks   []mark    `json:"marks,omitempty"`
}

type mark struct {
	Type  string `json:"type"`
	Attrs any    `json:"attrs"`
}

func unmarshalJSON(data []byte) (document, error) {
	var doc document
	err := json.Unmarshal(data, &doc)
	if err != nil {
		return document{}, err
	}
	return doc, nil
}

func contentToMarkdown(content []content) string {
	markdown := ""
	for _, c := range content {
		if c.Type == "embedly" || c.Type == "twitter" {
			continue
		}
		switch c.Type {
		case "heading":
			markdown += headingToMarkdown(c)
			markdown += "\n\n"
		case "paragraph":
			markdown += paragraphToMarkdown(c)
			markdown += "\n\n"
		case "text":
			markdown += textToMarkdown(c)
		case "image":
			markdown += imageToMarkdown(c)
			markdown += "\n\n"
		case "figure":
			markdown += figureToMarkdown(c)
		case "horizontalRule":
			markdown += hrToMarkdown(c)
			markdown += "\n\n"
		case "orderedList":
			markdown += orderedListToMarkdown(c)
		case "unorderedList":
			markdown += unorderedListToMarkdown(c)
		}
	}
	return markdown
}

func headingToMarkdown(content content) string {
	if content.Type != "heading" {
		return ""
	}
	level, ok := content.Attrs.(map[string]interface{})["level"].(float64)
	if !ok {
		level = 1
	}
	hashes := strings.Repeat("#", int(level))
	return hashes + " " + contentToMarkdown(content.Content)
}

func paragraphToMarkdown(content content) string {
	if content.Type != "paragraph" {
		return ""
	}
	return contentToMarkdown(content.Content)
}

func textToMarkdown(content content) string {
	if content.Type != "text" {
		return ""
	}
	text := content.Text
	for _, mark := range content.Marks {
		switch mark.Type {
		case "link":
			url, ok := mark.Attrs.(map[string]any)["href"].(string)
			if ok {
				text = fmt.Sprintf("[%s](%s)", text, url)
			}
		case "bold":
			text = fmt.Sprintf("**%s**", text)
		case "italic":
			text = fmt.Sprintf("_%s_", text)
		case "code":
			text = fmt.Sprintf("`%s`", text)
		case "strikethrough":
			text = fmt.Sprintf("~~%s~~", text)
		}
	}
	return text
}

func imageToMarkdown(content content) string {
	if content.Type != "image" {
		return ""
	}
	url, ok := content.Attrs.(map[string]any)["src"].(string)
	if !ok {
		return ""
	}
	return fmt.Sprintf("![%s](%s)", content.Text, url)
}

func figureToMarkdown(content content) string {
	if content.Type != "figure" {
		return ""
	}
	return contentToMarkdown(content.Content)
}

func hrToMarkdown(content content) string {
	if content.Type != "horizontalRule" {
		return ""
	}
	return "---"
}

func orderedListToMarkdown(content content) string {
	if content.Type != "orderedList" {
		return ""
	}
	start, ok := content.Attrs.(map[string]any)["start"].(float64)
	if !ok || start == 0 {
		start = 1
	}
	markdown := ""
	for _, c := range content.Content {
		markdown += listItemToMarkdown(c, fmt.Sprintf("%d.", int(start)))
		start++
	}
	return markdown
}

func unorderedListToMarkdown(content content) string {
	if content.Type != "unorderedList" {
		return ""
	}
	markdown := ""
	for _, c := range content.Content {
		markdown += listItemToMarkdown(c, "*")
	}
	return markdown
}

func listItemToMarkdown(content content, listMarker string) string {
	if content.Type != "listItem" {
		return ""
	}
	return listMarker + " " + contentToMarkdown(content.Content)
}
