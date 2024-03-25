import ts from "typescript";
import _ from "lodash";

export const getDecorators = (node: ts.HasDecorators) => {
    return _(ts.getDecorators(node)).filter(decorator => ts.isDecorator(decorator)).value();
}

export const getDecoratorValues = (decorator: ts.Decorator) => {
    const callExpression = decorator.expression as ts.CallExpression;
    const identifier = callExpression.expression as ts.Identifier;
    return {
        name: identifier.getText(),
        args: callExpression.arguments.filter(node => true) as ts.Expression[]
    }
}