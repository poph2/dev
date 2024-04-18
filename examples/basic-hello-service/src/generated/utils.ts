import {Person} from "../models/Person";
import {Database, getDb} from "../Main";
import {ReferenceExpression} from "kysely/dist/cjs/parser/reference-parser";
import {OperandValueExpressionOrList} from "kysely/dist/cjs/parser/binary-operation-parser";
import {ExtractTableAlias} from "kysely/dist/cjs/parser/table-parser";
import {Selectable} from "kysely";

// export const findByField = async (tableName: keyof Database & string, fieldKey: string, fieldValue: number): Promise<Person> => {
//     return await getDb<Database>().selectFrom(tableName)
//         .where('id', '=', fieldValue)
//         .selectAll().executeTakeFirstOrThrow();
// }

export const findPersonById = async (id: number): Promise<Person> => {
    return await getDb<Database>().selectFrom('person')
        .where('id', '=', id)
        .selectAll().executeTakeFirstOrThrow();
}
export const findByField = async <
    DB,
    TE extends keyof DB & string,
    TB extends ExtractTableAlias<DB, TE>,
    FK extends ReferenceExpression<DB, ExtractTableAlias<DB, TE>>,
    FV extends OperandValueExpressionOrList<DB, ExtractTableAlias<DB, TE>, FK>,
>(tableName: TE, fieldKey: FK, id: FV): Promise<Selectable<TB>> => {
    return await getDb<DB>().selectFrom(tableName)
        .where(fieldKey, '=', id)
        .selectAll().executeTakeFirstOrThrow() as unknown as Selectable<TB>;
}

// export const findPersionById = async (id: number): Promise<Person> => {
//     return findByField<Database, 'person', 'person', 'person', number>('person', 'id', id);
// }


type ValueOf<T> = T[keyof T];

export class BaseRepository<
    DB,
    // T
    // TB extends keyof DB,
    // TB extends keyof DB,
    // TE extends TableExpression<DB, keyof DB>,
    TE extends keyof DB & string
    // RE extends ReferenceExpression<DB, TB>,
    // VE extends OperandValueExpressionOrList<DB, TB, RE>,
> {

    constructor(protected readonly tableName: TE) {
    }

    get db() {
        return getDb<DB>()
    }

    async findByField<RE extends ReferenceExpression<DB, ExtractTableAlias<DB, TE>>,
        VE extends OperandValueExpressionOrList<DB, ExtractTableAlias<DB, TE>, RE>>(
        fieldKey: RE,
        fieldValue: VE): Promise<Selectable<ExtractTableAlias<DB, TE>>> {

        return await this.db.selectFrom(this.tableName)
            .where(fieldKey, '=', fieldValue)
            .selectAll().executeTakeFirstOrThrow() as unknown as Selectable<ExtractTableAlias<DB, TE>>;
    }

    async findByField2<RE extends ReferenceExpression<DB, ExtractTableAlias<DB, TE>>,
        VE extends OperandValueExpressionOrList<DB, ExtractTableAlias<DB, TE>, RE>>(
        fieldKey: RE,
        fieldValue: VE): Promise<Selectable<DB[TE]>> {
        return await this.db.selectFrom(this.tableName).where(fieldKey, '=', fieldValue)
            .selectAll().executeTakeFirstOrThrow() as unknown as Selectable<DB[TE]>
    }
}

export class PersonRepository {
    constructor() {
        // super('person');
    }

    get db() {
        return getDb<Database>()
    }

    // async findByField<
    //     RE extends ReferenceExpression<Database, ExtractTableAlias<Database, keyof Database & string>>,
    //     VE extends OperandValueExpressionOrList<DB, ExtractTableAlias<DB, TE>, RE>
    // >(fieldKey: RE, fieldValue: VE): Promise<Person> {
    //     return await this.db.selectFrom('person')
    //         .where(fieldKey, '=', fieldValue)
    //         .selectAll().executeTakeFirstOrThrow();
    // }

    // async findByFirstName(firstName: string): Promise<Person> {
    //     return await this.findByField('firstName', firstName);
    // }

    async findByFirstName2(firstName: string): Promise<Selectable<Database['person']>> {
        return this.db.selectFrom("person").where("firstName", "=", firstName)
            .selectAll().executeTakeFirstOrThrow()
    }
}
