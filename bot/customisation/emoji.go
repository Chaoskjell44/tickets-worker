package customisation

import (
	"fmt"

	"github.com/rxdn/gdl/objects"
	"github.com/rxdn/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1349851244050645083, false)
	EmojiOpen       = NewCustomEmoji("open", 1349851258961399881, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1349851274694103192, false)
	EmojiClose      = NewCustomEmoji("close", 1349851202707263519, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1349851216338878505, false)
	EmojiReason     = NewCustomEmoji("reason", 1349851336463745086, false)
	EmojiSubject    = NewCustomEmoji("subject", 1349851366025203743, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1349851401760805064, false)
	EmojiClaim      = NewCustomEmoji("claim", 1349851188614402148, false)
	EmojiPanel      = NewCustomEmoji("panel", 1349851292092334251, false)
	EmojiRating     = NewCustomEmoji("rating", 1349851321121112235, false)
	EmojiStaff      = NewCustomEmoji("staff", 1349851350720057356, false)
	EmojiThread     = NewCustomEmoji("thread", 1349851386904445009, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1349851171078017134, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1349851305577025587, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1349851229903388793, false)
	//EmojiTime       = NewCustomEmoji("time", 974006684622159952, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}
