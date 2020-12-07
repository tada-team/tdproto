package tdmarkup

import "testing"

func TestScanner(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		s := NewScanner("")

		if s.Position() != 0 {
			t.Error("pos must be 0")
		}

		if s.Rest() != 0 {
			t.Error("must be end")
		}

		if s.TakeNext() != EOF {
			t.Error("pop result must be EOF")
		}
	})

	t.Run("one", func(t *testing.T) {
		r := '客'
		s := NewScanner(string(r))

		if s.Rest() != 1 {
			t.Error("must be not end")
		}

		if s.Position() != 0 {
			t.Error("Index() must be 0")
		}

		if s.Next() != r {
			t.Errorf("Next() must be %s got %s at %d", string(r), string(s.Next()), s.Position())
		}

		if s.TakeNext() != r {
			t.Errorf("Pop() must be %s got %s at %d", string(r), string(s.Next()), s.Position())
		}

		if s.Rest() != 0 {
			t.Error("must be end")
		}

		if s.Next() != EOF {
			t.Errorf("Next() must be EOF got %s at %d", string(s.Next()), s.Position())
		}
	})

	t.Run("text", func(t *testing.T) {
		s := NewScanner("123")

		if ch := s.TakeNext(); ch != '1' {
			t.Error("next result must be 1")
		}
		if ch := s.Prev(); ch != EOF {
			t.Error("prev must be eof, got:", string(ch))
		}
		if ch := s.Next(); ch != '2' {
			t.Error("prev must be 2, got:", string(ch))
		}

		if ch := s.TakeNext(); ch != '2' {
			t.Error("next result must be 1")
		}
		if ch := s.Prev(); ch != '1' {
			t.Error("prev must be 1, got:", string(ch))
		}
		if ch := s.Next(); ch != '3' {
			t.Error("prev must be 3, got:", string(ch))
		}

		if ch := s.TakeNext(); ch != '3' {
			t.Error("next result must be 1")
		}
		if ch := s.Prev(); ch != '2' {
			t.Error("prev must be 2, got:", string(ch))
		}
		if ch := s.Next(); ch != EOF {
			t.Error("prev must be EOF, got:", string(ch))
		}
	})
}
