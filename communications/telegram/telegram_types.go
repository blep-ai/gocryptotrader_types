package telegram

type User struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
	Result      struct {
		ID           int64  `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		UserName     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"result"`
}
type GetUpdateResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
	Result      []struct {
		UpdateID           int64       `json:"update_id"`
		Message            MessageType `json:"message"`
		EditedMessage      interface{} `json:"edited_message"`
		ChannelPost        interface{} `json:"channel_post"`
		EditedChannelPost  interface{} `json:"edited_channel_post"`
		InlineQuery        interface{} `json:"inline_query"`
		ChosenInlineResult interface{} `json:"chosen_inline_result"`
		CallbackQuery      interface{} `json:"callback_query"`
		ShippingQuery      interface{} `json:"shipping_query"`
		PreCheckoutQuery   interface{} `json:"pre_checkout_query"`
	} `json:"result"`
}
type Message struct {
	Ok          bool        `json:"ok"`
	Description string      `json:"description"`
	Result      MessageType `json:"result"`
}
type MessageType struct {
	MessageID            int64               `json:"message_id"`
	From                 UserType            `json:"from"`
	Date                 int64               `json:"date"`
	Chat                 ChatType            `json:"chat"`
	ForwardFrom          UserType            `json:"forward_from"`
	ForwardFromChat      ChatType            `json:"forward_from_chat"`
	ForwardFromMessageID int64               `json:"forward_from_message_id"`
	ForwardSignature     string              `json:"forward_signature"`
	ForwardDate          int64               `json:"forward_date"`
	ReplyToMessage       interface{}         `json:"reply_to_message"`
	EditDate             int64               `json:"edit_date"`
	MediaGroupID         string              `json:"media_group_id"`
	AuthorSignature      string              `json:"author_signature"`
	Text                 string              `json:"text"`
	Entities             []MessageEntityType `json:"entities"`
	CaptionEntities      []MessageEntityType `json:"caption_entities"`
}
type MessageEntityType struct {
	Type   string   `json:"type"`
	Offset int64    `json:"offset"`
	Length int64    `json:"length"`
	URL    string   `json:"url"`
	User   UserType `json:"user"`
}
type UserType struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type ChatType struct {
	ID               int64  `json:"id"`
	Type             string `json:"type"`
	Title            string `json:"title"`
	UserName         string `json:"username"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	AllAdmin         bool   `json:"all_members_are_administrators"`
	Description      string `json:"description"`
	InviteLink       string `json:"invite_link"`
	StickerSetName   string `json:"sticker_set_name"`
	CanSetStickerSet bool   `json:"can_set_sticker_set"`
}
