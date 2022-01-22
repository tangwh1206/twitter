package common

import "strconv"

const (
	RedisKeyPrefixUid2User     = "twitter:uid:user"
	RedisKeyPrefixUsername2Uid = "twitter:username:uid"

	RedisKeyPrefixUid2SessionId = "twitter:uid:session_id"
	RedisKeyPrefixSessionId2uid = "twitter:session_id:uid"
)

func GenRedisKeyUidToUser(uid int64) string {
	return RedisKeyPrefixUid2User + strconv.FormatInt(uid, 10)
}

func GenRedisKeyUsernameToUid(username string) string {
	return RedisKeyPrefixUsername2Uid + username
}

func GenRedisKeyUid2SessionId(uid int64) string {
	return RedisKeyPrefixUid2SessionId + strconv.FormatInt(uid, 10)
}

func GenRedisKeySessionId2Uid(sessionId string) string {
	return RedisKeyPrefixSessionId2uid + sessionId
}
