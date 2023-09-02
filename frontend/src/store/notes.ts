import {defineStore} from "pinia";
import {
    CreateNote,
    DeleteNote,
    GetNote,
    GetNotes,
    RenameNote,
    UpdateNote
} from "../../wailsjs/go/main/App.js";
import {INote} from "../types";
import {appendNumberToTitle} from "../utils";

export const useNotesStore = defineStore('notes-store', {
    state: (): { nodes: INote[] } => ({
        nodes: []
    }),
    actions: {
        syncNote() {
            GetNotes().then((n) => {
                this.nodes = n ?? []
            });
        },
        async createNote(parentId: number) {
            const id = await CreateNote({
                id: 0,
                title: appendNumberToTitle("New note", this.nodes.map((n) => n.title)),
                body: "",
                isDir: 0,
                parentId: parentId,
                isFav: 0,
                createTimestamp: Date.now(),
                updateTimestamp: Date.now(),
            })
            if (id) this.syncNote()
            return id
        },
        createFolder(name: string) {
            CreateNote({
                id: 0,
                title: appendNumberToTitle(name, this.nodes.map((n) => n.title)),
                body: "",
                isDir: 1,
                parentId: 0,
                isFav: 0,
                createTimestamp: Date.now(),
                updateTimestamp: Date.now(),
            }).then(r => {
                if (r) this.syncNote()
            })
        },
        renameNote(id: number, title: string) {
            RenameNote(id, appendNumberToTitle(title, this.nodes.filter((n) => n.id !== id).map((n) => n.title))).then((r) => {
                if (r) this.syncNote()
            })
        },
        deleteNote(id: number) {
            DeleteNote(id).then((r) => {
                if (r) this.syncNote()
            })
        },
        getNote(id: number) {
            return GetNote(id);
        },
        getFolderNotes(id: number) {
            return this.nodes.filter((n) => n.parentId === id)
        },
        updateNote(id: number, body: string) {
            return UpdateNote(id, body);
        },


    },
})
