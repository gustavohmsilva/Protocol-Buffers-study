protoc          -I src/                 --go_out=src/           src/simple/simple.proto
protoc          -I src/                 --go_out=src/           src/enum/enum.proto
protoc          -i src/                 --go_out=src/           src/complex/complex.proto
#complier       source folder           output in folder        full path to file