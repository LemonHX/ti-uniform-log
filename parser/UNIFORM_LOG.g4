grammar UNIFORM_LOG;

logs: log ('\n' log)*;
log: date_time level location custom_message message*;

date_time: '[' DATE_TIME ']';
level: '[' LEVEL ']';
location: '[' '<unknown>' ']' | '[' known_location ']';
custom_message: '[' key ']';
message: '[' kv_pair ']';

known_location : PATHLOC ':' NUM;
kv_pair: key '=' value;
key: IDENT | STRING;
value: IDENT | STRING | NUM | DURATION | IP;

IP: [0-9]+'.'[0-9]+'.'[0-9]+'.'[0-9]+':'[0-9]+;
DATE_TIME:
	NNNN '/' NN '/' NN WS NN ':' NN ':' NN '.' NNN WS ZZZZZ;
NUM: [0-9]+;
FLOAT: NUM '.' NUM;
DURATION: (FLOAT ('s'|'m'|'h'))+;
LEVEL: [A-Z]+;
PATHLOC: ([a-zA-Z0-9\-_])+'.'([a-zA-Z0-9\-_])+;
STRING: '"' (ESC | SAFECODEPOINT)* '"';
IDENT: SAFEIDENTPOINT+;

fragment NNNN: [0-9][0-9][0-9][0-9];
fragment NN: [0-9][0-9];
fragment NNN: [0-9][0-9][0-9];
fragment ZZZZZ: ('+' | '-')? NN ':' NN;
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];
fragment SAFECODEPOINT: ~ ["\\\u0000-\u001F];
fragment SAFEIDENTPOINT:
	~[\u0000-\u0020\u0022\u003d\u005b\u005d:];
WS: [ \t\n\r]+ -> skip;