package push

type Results []Result

func (rs Results) AllOf() Result {
	return Result{}  // TODO
}

func (rs Results) AnyOf() Result {
	return Result{}  // TODO
}