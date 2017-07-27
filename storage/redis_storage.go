package storage

import (
	"github.com/garyburd/redigo/redis"
	"github.com/goribun/ProxyPool/util"
)

//使用连接池
var RedisClient = &redis.Pool{
	MaxIdle:   util.Cfg.Redis.MaxIdle,
	MaxActive: util.Cfg.Redis.MaxActive,
	Dial: func() (redis.Conn, error) {
		conn, err := redis.Dial("tcp", string(util.Cfg.Redis.Addr))
		if err != nil {
			return nil, err
		}
		return conn, nil
	},
}

//redis存储结构体
type RedisStorage struct {
	key string
}

//创建RedisStorage实例
func NewRedisStorage() *RedisStorage {
	return &RedisStorage{key: util.Cfg.Redis.Key}
}

//返回redis连接
func (s *RedisStorage) GetRedisConn() redis.Conn {
	return RedisClient.Get()
}

//向Redis集合增加IP
func (s *RedisStorage) Add(ip string) error {
	conn := s.GetRedisConn()
	defer conn.Close()
	_, err := conn.Do("SADD", s.key, ip)
	if err != nil {
		return err
	}
	return nil
}

//从Redis中随机取出一个IP
func (s *RedisStorage) GetOne() (string, error) {
	conn := s.GetRedisConn()
	defer conn.Close()
	ip, err := redis.String(conn.Do("SRANDMEMBER", s.key))

	if err != nil {
		return "", err
	}
	return ip, nil
}

// 统计redis集合中IP个数
func (s *RedisStorage) Count() int {
	conn := s.GetRedisConn()
	defer conn.Close()

	num, err := redis.Int(conn.Do("SCARD", s.key))
	if err != nil {
		num = 0
	}
	return num
}

// 删除redis集合中指定IP
func (s *RedisStorage) Delete(ip string) error {
	conn := s.GetRedisConn()
	defer conn.Close()
	_, err := conn.Do("SREM", s.key, ip)
	if err != nil {
		return err
	}
	return nil
}

//从redis集合中取得所有IP
func (s *RedisStorage) GetAll() ([]string, error) {
	conn := s.GetRedisConn()
	defer conn.Close()

	ips, err := redis.Strings(conn.Do("SMEMBERS", s.key))

	if err != nil {
		return nil, err
	}
	return ips, nil
}
