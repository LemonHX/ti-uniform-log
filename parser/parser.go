package parser

type LogHeaderSection struct {
	DateTime string    `@DateTime`
	LogLevel string    `@LogLevel`
	Location string    `@Location`
	Message  []Message `@@*`
}

type Message struct {
	KVPair  KVPair `@@`
	Message string `@Message`
	//Unquoted string `| @Unquoted`
}

type KVPair struct {
	Key   string `@Ident`
	Value Value  `@@`
}

type Key struct {
	String string `@String`
	Ident  string `| @Ident`
}

type Value struct {
	String string `@String`
	//IPAddr string ``
	//TimeDuration string ``
}

type FuckMe struct {
	Name string `@Ident`
}
