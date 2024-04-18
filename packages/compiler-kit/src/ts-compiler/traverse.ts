import ts from "typescript";

export type Aliases = {
  ClassDeclaration: ts.ClassDeclaration;
  MethodDeclaration: ts.MethodDeclaration;
  TypeAliasDeclaration: ts.TypeAliasDeclaration;
  InterfaceDeclaration: ts.InterfaceDeclaration;
}

export type TraverseOptions = {
  [key in keyof Aliases]?: (node: Aliases[key]) => void;
} & {
  depth?: number;
}

export const traverse = (node: ts.Node, opts: TraverseOptions) => {
  const depth = opts.depth || 10;
  if(opts.ClassDeclaration && ts.isClassDeclaration(node)) {
    opts.ClassDeclaration(node);
  }
  else if(opts.MethodDeclaration && ts.isMethodDeclaration(node)) {
    opts.MethodDeclaration(node);
  }
  else if(opts.TypeAliasDeclaration && ts.isTypeAliasDeclaration(node)) {
    opts.TypeAliasDeclaration(node);
  }
  else if(opts.InterfaceDeclaration && ts.isInterfaceDeclaration(node)) {
    opts.InterfaceDeclaration(node);
  }
  if(depth > 0) {
    ts.forEachChild(node, child => traverse(child, {...opts, depth: depth - 1}));
  }
}