import ts from "typescript";
import { traverse } from "./traverse";
import fs from "fs";


// type Aliases = {
//     ClassDeclaration: ts.ClassDeclaration;
//     MethodDeclaration: ts.MethodDeclaration;
// }
//
// type TraverseOptions = {
//     [key in keyof Aliases]?: (node: Aliases[key]) => void;
// } & {
//     depth?: number;
// }



export const extractClasses = (node: ts.SourceFile, criteriaFn?: (n: ts.ClassDeclaration) => boolean, depth: number = 1) => {
    const classes: ts.ClassDeclaration[] = [];

    traverse(node, {
        ClassDeclaration: (_node) => {
            if(!criteriaFn || criteriaFn(_node)) {
                classes.push(_node);
            }
        },
        depth
    });

    return classes;

}

export const extractMethods = (node: ts.Node, criteriaFn?: (n: ts.MethodDeclaration) => boolean, depth: number = 1) => {
    const methods: ts.MethodDeclaration[] = [];

    traverse(node, {
        MethodDeclaration: (_node) => {
            if(!criteriaFn || criteriaFn(_node)) {
                methods.push(_node);
            }
        },
        depth
    })

    return methods;
}