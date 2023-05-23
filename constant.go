package tempest

import "fmt"

const (
	DISCORD_API_URL  = "https://discord.com/api/v10"
	DISCORD_CDN_URL  = "https://cdn.discordapp.com"
	USER_AGENT       = "DiscordApp https://github.com/Amatsagu/tempest"
	EPOCH            = 1420070400000 // Discord epoch in milliseconds
	ROOT_PLACEHOLDER = "-"
)

// Prepare those replies as they never change so there's no point in re-creating them each time.
var (
	PING_RESPONSE_RAW_BODY            = []byte(fmt.Sprintf(`{"type":%d}`, PONG_RESPONSE_TYPE))
	ACKNOWLEDGE_RESPONSE_RAW_BODY     = []byte(fmt.Sprintf(`{"type":%d}`, DEFERRED_UPDATE_MESSAGE_RESPONSE_TYPE))
	UNKNOWN_COMMAND_RESPONSE_RAW_BODY = []byte(fmt.Sprintf(`{"type":%d,"data":{"content":"Oh snap! It looks like you tried to trigger (/) unknown command. Please report this bug to bot owner.","flags":64}}`, CHANNEL_MESSAGE_WITH_SOURCE_RESPONSE_TYPE))
)
