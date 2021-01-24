package main

// Dictionary は definition と word を保持するもの
type Dictionary map[string]string

var (
	// ErrNotFound は definition が与えられた word を見つけることができなかった時のエラー
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists は既にある word を追加しようとしてる時のエラー
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExist は Update しようとするもその定義自体が存在しなかった時のエラー
	ErrWordDoesNotExist = DictionaryErr("cannot update because it does not exist")
)

// DictionaryErr は dictionary を処理する際に起こりうるエラーをまとめたもの
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary の中の word を Search して見つけるためのメソッド
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Dictionary に definition と word を Add するためのメソッド
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
func (d Dictionary) Update(word, definition string) error {
	//d[word] = definition
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// 引数に与えられた word が Dictionary にあるとその word と 値を Delete する
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
