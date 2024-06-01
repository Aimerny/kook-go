package model

import "time"

// Card Message
type CardMessageReq []*CardMessage

type CardThemeType string
type CardSize string

const (
	ThemeTypeSecondary CardThemeType = "secondary"
	ThemeTypePrimary   CardThemeType = "primary"
	ThemeTypeSuccess   CardThemeType = "success"
	ThemeTypeDanger    CardThemeType = "danger"
	ThemeTypeWarning   CardThemeType = "warning"
	ThemeTypeInfo      CardThemeType = "info"
	ThemeTypeNone      CardThemeType = "none"

	SizeXs CardSize = "xs"
	SizeSm CardSize = "sm"
	SizeMd CardSize = "md"
	SizeLg CardSize = "lg"
)

type CardModule struct {
	Type      string        `json:"type"`
	Theme     CardThemeType `json:"theme,omitempty"`
	Text      CardText      `json:"text,omitempty"`
	Size      CardSize      `json:"size,omitempty"`
	Modules   []CardModule  `json:"modules,omitempty"`
	Fields    []CardModule  `json:"fields,omitempty"`
	Elements  []CardModule  `json:"elements,omitempty"`
	Mode      string        `json:"mode,omitempty"`
	Accessory *CardModule   `json:"accessory,omitempty"`
	Value     string        `json:"value,omitempty"`
	Src       string        `json:"src,omitempty"`
	StartTime int64         `json:"startTime,omitempty"`
	EndTime   int64         `json:"endTime,omitempty"`
}

type CardText struct {
	Type    string     `json:"type,omitempty"`
	Content string     `json:"content,omitempty"`
	Cols    int        `json:"cols,omitempty"`
	Fields  []CardText `json:"fields,omitempty"`
}

type CardMessage CardModule

func NewCard(theme CardThemeType, size CardSize) *CardModule {
	return &CardModule{
		Type:  "card",
		Theme: theme,
		Size:  size,
	}
}

func NewTextHeader(content string) *CardModule {
	return &CardModule{
		Type: "header",
		Text: CardText{
			Type:    "plain-text",
			Content: content,
		},
	}
}

func NewText(content string) *CardModule {
	return &CardModule{
		Type: "section",
		Text: CardText{
			Type:    "plain-text",
			Content: content,
		},
	}
}

func NewKMarkdown(content string) *CardModule {
	return &CardModule{
		Type: "section",
		Text: CardText{
			Type:    "kmarkdown",
			Content: content,
		},
	}
}

func NewMultiLinesKMarkdown(lines []string) *CardModule {
	fields := make([]CardText, len(lines))
	for _, line := range lines {
		fields = append(fields, CardText{
			Type:    "kmarkdown",
			Content: line,
		})
	}
	return &CardModule{
		Type: "section",
		Text: CardText{
			Type:   "paragraph",
			Cols:   len(lines),
			Fields: fields,
		},
	}
}

func NewRightImgKMarkdown(content, src string) *CardModule {
	module := NewKMarkdown(content)
	module.Mode = "right"
	module.Accessory = &CardModule{
		Type: "image",
		Src:  src,
		Size: SizeLg,
	}
	return module
}

func NewImage(src string) *CardModule {
	return &CardModule{
		Type: "container",
		Elements: []CardModule{
			{
				Type: "image",
				Src:  src,
			},
		},
	}
}

func NewImageGroup(src []string) *CardModule {
	elements := make([]CardModule, len(src))
	for _, line := range src {
		elements = append(elements, CardModule{
			Type: "image",
			Src:  line,
		})
	}
	return &CardModule{
		Type:     "image-group",
		Elements: elements,
	}
}

func NewDivider() *CardModule {
	return &CardModule{
		Type: "divider",
	}
}

func NewComment(elements []CardModule) *CardModule {
	return &CardModule{
		Type:     "context",
		Elements: elements,
	}
}

func NewCountdownDay(endTime time.Time) *CardModule {
	return &CardModule{
		Type:    "countdown",
		Mode:    "day",
		EndTime: endTime.UnixMilli(),
	}
}

func NewCountdownHour(endTime *time.Time) *CardModule {
	return &CardModule{
		Type:    "countdown",
		Mode:    "hour",
		EndTime: endTime.UnixMilli(),
	}
}

func NewCountdownSecond(startTime, endTime *time.Time) *CardModule {
	return &CardModule{
		Type:      "countdown",
		Mode:      "second",
		StartTime: startTime.UnixMilli(),
		EndTime:   endTime.UnixMilli(),
	}
}
