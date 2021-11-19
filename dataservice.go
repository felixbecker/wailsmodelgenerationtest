package main

type Service struct {
}

type Todo struct {
	Description string
}

func (s *Service) All() []Todo {

	return []Todo{
		{Description: "hello world"},
	}
}
