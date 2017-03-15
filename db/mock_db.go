package db

type MockSdb struct {
	OnFindGitHubUserByName func(string) (*GitHubUser, error)
	OnSaveRecord           func(interface{}) error
}

func (mdb *MockSdb) FindGitHubUserByName(name string) (*GitHubUser, error) {
	return mdb.OnFindGitHubUserByName(name)
}

func (mdb *MockSdb) SaveRecord(record interface{}) error {
	return mdb.OnSaveRecord(record)
}
