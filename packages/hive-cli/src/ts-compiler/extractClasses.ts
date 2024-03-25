import ts from 'typescript';
import * as fs from "fs";

type Aliases = {
    ClassDeclaration: ts.ClassDeclaration;
    MethodDeclaration: ts.MethodDeclaration;
}

type TraverseOptions = {
    [key in keyof Aliases]?: (node: Aliases[key]) => void;
} & {
    depth?: number;
}


const traverse = (node: ts.Node, opts: TraverseOptions) => {
    const depth = opts.depth || 10;
    if(opts.ClassDeclaration && ts.isClassDeclaration(node)) {
        opts.ClassDeclaration(node);
    }
    else if(opts.MethodDeclaration && ts.isMethodDeclaration(node)) {
        opts.MethodDeclaration(node);
    }
    if(depth > 0) {
        ts.forEachChild(node, child => traverse(child, {...opts, depth: depth - 1}));
    }
}

export const readSourceFile = (filePath: string): ts.SourceFile => {
    const sourceCode = fs.readFileSync(filePath, 'utf-8');

    return ts.createSourceFile(
        filePath,
        sourceCode,
        ts.ScriptTarget.Latest,
        /* setParentNodes */ true
    );
}

export const extractClasses = (sourceFile: ts.SourceFile, criteriaFn?: (node: ts.ClassDeclaration) => boolean, depth: number = 1) => {
    const classes: ts.ClassDeclaration[] = [];

    traverse(sourceFile, {
        ClassDeclaration: (node) => {
            if(!criteriaFn || criteriaFn(node)) {
                classes.push(node);
            }
        },
        depth
    });

    return classes;

}

export const extractMethods = (node: ts.Node, criteriaFn?: (node: ts.MethodDeclaration) => boolean, depth: number = 1) => {
    const methods: ts.MethodDeclaration[] = [];
    traverse(node, {
        MethodDeclaration: (node) => {
            if(!criteriaFn || criteriaFn(node)) {
                methods.push(node);
            }
        },
        depth
    })
    return methods;
}