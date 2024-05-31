package convert

import "testing"

func TestPostJSONToMarkdown(t *testing.T) {
	tests := []struct {
		name string
		json []byte
		want string
	}{
		{
			name: "text",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "text",
						"text": "Hello, world!"
					}
				]
			}`),
			want: "Hello, world!",
		},
		{
			name: "paragraph",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "paragraph",
						"content":[
							{
								"type": "text",
								"text": "Hello, world!"
							}
						]
					}
				]
			}`),
			want: "Hello, world!\n\n",
		},
		{
			name: "image",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "image",
						"attrs": {"src": "https://example.com/image.jpg"},
						"text": "Image"
					}
				]
			}`),
			want: "![Image](https://example.com/image.jpg)\n\n",
		},
		{
			name: "figure",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "figure",
						"content":[
							{
								"type": "image",
								"attrs": {"src": "https://example.com/image.jpg"},
								"text": "Image"
							}
						]
					}
				]
			}`),
			want: "![Image](https://example.com/image.jpg)\n\n",
		},
		{
			name: "orderedList",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "orderedList",
						"attrs": {"start": 20},
						"content":[
							{
								"type": "listItem",
								"content": [
									{
										"type": "paragraph",
										"content":[
											{
												"type": "text",
												"text": "Hello, world!"
											}
										]
									}
								]
							},
							{
								"type": "listItem",
								"content": [
									{
										"type": "paragraph",
										"content":[
											{
												"type": "text",
												"text": "Hello, world 2!"
											}
										]
									}
								]
							}
						]
					}
				]
			}`),
			want: "20. Hello, world!\n\n21. Hello, world 2!\n\n",
		},
		{
			name: "unorderedList",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "unorderedList",
						"content":[
							{
								"type": "listItem",
								"content": [
									{
										"type": "paragraph",
										"content":[
											{
												"type": "text",
												"text": "Hello, world!"
											}
										]
									}
								]
							},
							{
								"type": "listItem",
								"content": [
									{
										"type": "paragraph",
										"content":[
											{
												"type": "text",
												"text": "Hello, world 2!"
											}
										]
									}
								]
							}
						]
					}
				]
			}`),
			want: "* Hello, world!\n\n* Hello, world 2!\n\n",
		},
		{
			name: "strikethrough",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "text",
						"text": "Hello, world!",
						"marks": [
							{
								"type": "strikethrough"
							}
						]
					}
				]
			}`),
			want: "~~Hello, world!~~",
		},
		{
			name: "horizontalRule",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "horizontalRule"
					}
				]
			}`),
			want: "---\n\n",
		},
		{
			name: "code",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "text",
						"text": "Hello, world!",
						"marks": [
							{
								"type": "code"
							}
						]
					}
				]
			}`),
			want: "`Hello, world!`",
		},
		{
			name: "link",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "text",
						"text": "Hello, world!",
						"marks": [
							{
								"type": "link",
								"attrs": {"href": "https://example.com"}
							}
						]
					}
				]
			}`),
			want: "[Hello, world!](https://example.com)",
		},
		{
			name: "text_multimark",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "text",
						"text": "Hello, world!",
						"marks": [
							{
								"type": "link",
								"attrs": {"href": "https://example.com"}
							},
							{
								"type": "bold"
							},
							{
								"type": "italic"
							}
						]
					}
				]
			}`),
			want: "_**[Hello, world!](https://example.com)**_",
		},
		{
			name: "heading",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "heading",
						"attrs": {"level": 1},
						"content":[
							{
								"type": "text",
								"text": "Hello, world!"
							}
						]
					}
				]
			}`),
			want: "# Hello, world!\n\n",
		},
		{
			name: "heading",
			json: []byte(`{
				"type":"doc",
				"content":[
					{
						"type": "heading",
						"attrs": {"level": 2},
						"content":[
							{
								"type": "text",
								"text": "Hello, world!"
							}
						]
					}
				]
			}`),
			want: "## Hello, world!\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostJSONToMarkdown(tt.json)
			if err != nil {
				t.Errorf("PostJSONToMarkdown() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("PostJSONToMarkdown() = %q want %q", got, tt.want)
			}
		})
	}
}
