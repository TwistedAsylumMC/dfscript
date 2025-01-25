declare namespace sql {
    interface Database {
        begin(): Transaction;
        close(): void;
        exec(query: string, ...args: any[]): ExecResult;
        ping(): void;
        prepare(query: string): Statement
        query(query: string, ...args: any[]): object[];
        queryRow(query: string, ...args: any[]): object;
        setConnMaxIdleTime(d: number): void;
        setConnMaxLifetime(d: number): void;
        setMaxIdleConns(n: number): void;
        setMaxOpenConns(n: number): void;
        stats(): Stats;
    }
    function open(driver: string, dsn: string): Database;

    interface ExecResult {
        lastInsertId(): number;
        rowsAffected(): number;
    }

    interface Statement {
        close(): void;
        exec(...args: any[]): ExecResult;
        query(...args: any[]): object[];
        queryRow(...args: any[]): object;
    }

    interface Stats {
        maxOpenConnections: number;

        openConnections: number;
        inUse: number;
        idle: number;

        waitCount: number;
        waitDuration: number;
        maxIdleClosed: number;
        maxIdleTimeClosed: number;
        maxLifetimeClosed: number;
    }

    interface Transaction {
        commit(): void;
        exec(query: string, ...args: any[]): ExecResult;
        prepare(query: string): Statement;
        query(query: string, ...args: any[]): object[];
        queryRow(query: string, ...args: any[]): object;
        rollback(): void;
        stmt(statement: Statement): Statement;
    }
}