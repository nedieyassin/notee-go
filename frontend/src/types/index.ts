export interface INote {
    id: number,
    title: string,
    body: string,
    isDir: number,
    parentId: number,
    isFav: number,
    createTimestamp: number,
    updateTimestamp: number,
}