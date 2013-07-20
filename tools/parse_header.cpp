#include "clang/AST/ASTContext.h"
#include "clang/AST/ASTConsumer.h"
#include "clang/AST/RecursiveASTVisitor.h"
#include "clang/Frontend/CompilerInstance.h"
#include "clang/Frontend/FrontendAction.h"
#include "clang/Tooling/Tooling.h"

#include <iostream>
#include <fstream>
#include <string>
#include <sstream>

using namespace std;
using namespace clang;

class FindNamedClassVisitor
: public RecursiveASTVisitor<FindNamedClassVisitor> {
  public:
    explicit FindNamedClassVisitor(ASTContext *Context)
      : Context(Context) {}

    bool VisitFunctionDecl(FunctionDecl *Declaration) {
      cout << "func|";
      auto name_info = Declaration->getNameInfo().getName();
      auto name = name_info.getAsString();
      cout << name;
      auto ret_type = Declaration->getResultType();
      auto ret_type_str = ret_type.getAsString();
      cout << "|" << ret_type_str;
      for (auto pi = Declaration->param_begin(), end = Declaration->param_end();
          pi != end; pi++ ) {
        auto param_type = (*pi)->getOriginalType();
        auto param_type_str =param_type.getAsString();
        auto name_id = (*pi)->getIdentifier();
        cout << "|" << param_type_str;
        if (name_id != NULL) {
          cout << "@" << name_id->getName().str();
        }
      }
      cout << endl;
      return true;
    }

    bool VisitEnumDecl(EnumDecl *Declaration) {
      return true;
    }

  private:
    ASTContext *Context;
};

class FindNamedClassConsumer : public clang::ASTConsumer {
  public:
    explicit FindNamedClassConsumer(ASTContext *Context)
      : Visitor(Context) {}

    virtual void HandleTranslationUnit(clang::ASTContext &Context) {
      Visitor.TraverseDecl(Context.getTranslationUnitDecl());
    }
  private:
    FindNamedClassVisitor Visitor;
};

class FindNamedClassAction : public clang::ASTFrontendAction {
  public:
    virtual clang::ASTConsumer *CreateASTConsumer(
        clang::CompilerInstance &Compiler, llvm::StringRef InFile) {
      return new FindNamedClassConsumer(&Compiler.getASTContext());
    }
};

int main(int argc, char **argv) {
  if (argc > 1) {
    std::ifstream input(argv[1]);
    std::string content((std::istreambuf_iterator<char>(input)),
        (std::istreambuf_iterator<char>()));
    std::vector<std::string> args;
    for (int i = 2; i < argc; i++) {
      std::istringstream iss(argv[i]);
      std::string token;
      while (iss >> token) {
        args.push_back(token);
      }
    }
    clang::tooling::runToolOnCodeWithArgs(new FindNamedClassAction, content, args, argv[1]);
  }
}

