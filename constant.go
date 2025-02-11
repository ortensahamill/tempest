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
	private_PING_RESPONSE_RAW_BODY            = []byte(fmt.Sprintf(`{"type":%d}`, PONG_RESPONSE_TYPE))
	private_ACKNOWLEDGE_RESPONSE_RAW_BODY     = []byte(fmt.Sprintf(`{"type":%d}`, DEFERRED_UPDATE_MESSAGE_RESPONSE_TYPE))
	private_UNKNOWN_COMMAND_RESPONSE_RAW_BODY = []byte(fmt.Sprintf(`{"type":%d,"data":{"content":"Oh uh.. It looks like you tried to trigger (/) unknown command. Please report this bug to bot owner.","flags":64}}`, CHANNEL_MESSAGE_WITH_SOURCE_RESPONSE_TYPE))
)

// Those are used to replace seemingly empty slice into empty array after marsalling struct to json string.
var (
	private_REST_NULL_SLICE_FIND    []byte = []byte("[null]")
	private_REST_NULL_SLICE_REPLACE []byte = []byte("[]")
)
