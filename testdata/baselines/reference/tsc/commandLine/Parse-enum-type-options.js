
currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::--moduleResolution nodenext  first.ts --module nodenext --target esnext --moduleDetection auto --jsx react --newLine crlf

ExitStatus:: 0

ParsedCommandLine::{
    "parsedConfig": {
        "compilerOptions": {
            "jsx": 3,
            "module": 199,
            "moduleResolution": 99,
            "moduleDetection": 1,
            "newLine": 1,
            "target": 99
        },
        "watchOptions": {
            "watchInterval": null,
            "watchFile": 0,
            "watchDirectory": 0,
            "fallbackPolling": 0,
            "synchronousWatchDirectory": null,
            "excludeDirectories": null,
            "excludeFiles": null
        },
        "typeAcquisition": null,
        "fileNames": [
            "first.ts"
        ],
        "projectReferences": null
    },
    "configFile": null,
    "errors": [],
    "raw": {
        "moduleResolution": 99,
        "module": 199,
        "target": 99,
        "moduleDetection": 1,
        "jsx": 3,
        "newLine": 1
    },
    "compileOnSave": null
}
Output::
//// [/home/src/tslibs/TS/Lib/lib.esnext.full.d.ts] *Lib*
/// <reference no-default-lib="true"/>
interface Boolean {}
interface Function {}
interface CallableFunction {}
interface NewableFunction {}
interface IArguments {}
interface Number { toExponential: any; }
interface Object {}
interface RegExp {}
interface String { charAt: any; }
interface Array<T> { length: number; [n: number]: T; }
interface ReadonlyArray<T> {}
interface SymbolConstructor {
    (desc?: string | number): symbol;
    for(name: string): symbol;
    readonly toStringTag: symbol;
}
declare var Symbol: SymbolConstructor;
interface Symbol {
    readonly [Symbol.toStringTag]: string;
}
declare const console: { log(msg: any): void; };

