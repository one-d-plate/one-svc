package pkg

func GetOffset(page int) int {
	const pageSize = 10
	if page < 1 {
		page = 1
	}
	return (page - 1) * pageSize
}
