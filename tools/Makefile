CXX := clang++
RTTIFLAG := -fno-rtti
LLVMCONFIG := llvm-config

CXXFLAGS := $(shell $(LLVMCONFIG) --cxxflags) $(RTTIFLAG) --std=c++11
LLVMLDFLAGS := $(shell $(LLVMCONFIG) --ldflags --libs $(LLVMCOMPONENTS))

SOURCES = parse_header.cpp

OBJECTS = $(SOURCES:.cpp=.o)
EXES = $(OBJECTS:.o=)
	CLANGLIBS = \
							-lclangTooling\
							-lclangFrontendTool\
							-lclangFrontend\
							-lclangDriver\
							-lclangSerialization\
							-lclangCodeGen\
							-lclangParse\
							-lclangSema\
							-lclangStaticAnalyzerFrontend\
							-lclangStaticAnalyzerCheckers\
							-lclangStaticAnalyzerCore\
							-lclangAnalysis\
							-lclangARCMigrate\
							-lclangRewriteFrontend\
							-lclangRewriteCore\
							-lclangEdit\
							-lclangAST\
							-lclangLex\
							-lclangBasic\
							$(shell llvm-config --libs)

all: $(OBJECTS) $(EXES)
	./parse_header all.h `pkg-config --cflags-only-I elementary emotion eio` > header_info

%: %.o
	$(CXX) -o $@ $< $(CLANGLIBS) $(LLVMLDFLAGS)

clean:
	-rm -f $(EXES) $(OBJECTS) *~
