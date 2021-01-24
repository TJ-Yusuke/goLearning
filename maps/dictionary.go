package main

// Dictionary は definition と word を保持するもの
type Dictionary map[string]string

var (
	// ErrNotFound は definition が与えられた word を見つけることができなかったという意味
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists は既にある word を追加しようとしていますという意味
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

// DictionaryErr は dictionary を処理する際に起こりうるエラーをまとめたもの
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// dictionary の中の word を Search して見つけるためのメソッド
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// dictionary に definition と word を Add するためのメソッド
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// 既にある定義を与えられた word で Update するためのメソッド
func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
