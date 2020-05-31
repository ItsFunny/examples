package authentication

import (
	"errors"
	"myLibrary/go-library/go/converters"
)

// 权限校验
// 考虑到可能后期权限数量大大增加,因而采用bitset的形式
// 缺点: 无法进行权限的快速累加,只可遍历添加
// 优点: 内存放的下多少个8字节的数据,就能存储多少个权限值,支持动态添加

type Authority []uint64

func NewAuthority() Authority {
	container := make(Authority, 1)
	return container
}

// 通过大端,将数据,从[]uint64 转换为字节数组
func (receiver Authority) BigEndianConvt2Bytes() []byte {
	if len(receiver) == 0 {
		return make([]byte, 8)
	}
	bs := make([]byte, 0)
	for _, v := range receiver {
		bs = append(bs, converter.BigEndianUInt642Bytes(v)...)
	}
	return bs
}

// 8个为一位
func BigEndianConvtBytes2Authority(bytes []byte) (Authority, error) {
	l := len(bytes)

	if l%8 != 0 {
		return nil, errors.New("参数错误,参数必须为8的整数倍")
	}
	valueCounts := l / 8
	authority := make([]uint64, valueCounts)
	count := 0
	i := 0
	for {
		if count >= valueCounts {
			break
		}
		// authority.AddAuthentication(AuthValue(converter.BigEndianBytes2Int64(bytes[i : i+8])))
		// authority[i] = 1

		authority[count]=converter.BigEndianBytes2Int64(bytes[i:i+8])
		// bs := bytes[i : i+8]
		// v := converter.BigEndianBytes2Int64(bs)
		// fmt.Println(v)
		// authority = append(authority, v)
		i += 8
		count++
	}

	return authority, nil
}

func (receiver Authority) AddAuthentication(value AuthValue) Authority {
	t := receiver
	if t == nil {
		t = NewAuthority()
	}
	i := value >> 6
	p := value & 0x3f

	if i > AuthValue(len(t)-1) {
		tt := make(Authority, i+1)
		copy(tt, t)
		t = tt
	}
	t[i] |= 1 << p

	return t
}

func (receiver Authority) CheckAuthentication(authValues ...AuthValue) bool {
	for _, v := range authValues {
		i := v >> 6
		p := v & 0x3f
		if i > AuthValue(len(receiver)) {
			return false
		}
		if (receiver[i] & (1 << p)) == 0 {
			return false
		}
	}
	return true
	// i := authValue >> 6
	// p := authValue & 0x3f
	// if i > AuthValue(len(receiver)) {
	// 	return false
	// }
	// return (receiver[i] & (1 << p)) != 0
}

func (receiver Authority) CheckAuthority(authority Authority) bool {
	l1 := len(receiver)
	l2 := len(authority)
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if receiver[i]&authority[i] == 0 {
			return false
		}
	}
	return true
}

func (receiver Authority) DeleteAuthentication(authValue AuthValue) Authority {
	t := receiver
	i := authValue >> 6
	p := authValue & 0x3f

	if i > AuthValue(len(t)) {
		return t
	}
	t[i] &= ^(1 << p)

	return t
}

func (receiver Authority) CreateSuperSuperAdmin() Authority {
	t := receiver
	for i := 0; i < 64; i++ {
		t = t.AddAuthentication(AuthValue(i))
	}
	return t
}
