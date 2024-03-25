import ts from "typescript";

const getDecorators = (node: ts.MethodDeclaration) => {
    return ts.getDecorators(node);
}