<script setup lang="ts">

import {StickyNote, Pencil, Trash2} from "lucide-vue-next";
import {PropType} from "vue";
import {INote} from "../../types";
import {useNotesStore} from "../../store/notes";
import {confirmDialog, promptDialog} from "../dialogs";
import {formatDateString} from "../../utils";

defineProps({
  note: {
    type: Object as PropType<INote>,
    required: true
  },
  fromFolder: Number
})

const notesStore = useNotesStore();

const deleteNote = (id: number) => {
  confirmDialog({
    title: 'Delete note?',
    message: 'Are you sure you want to delete this note?',
    action: 'Yes, Delete'
  }).then((res) => {
    if (res) {
      notesStore.deleteNote(id)
    }
  })
}

const renameNote = (id: number, title: string) => {
  promptDialog({
    title: 'Edit note title',
    placeholder: 'Enter note title',
    action: 'Edit',
    value: title
  }).then((res) => {
    if (res) {
      notesStore.renameNote(id, res);
    }
  })
}
</script>

<template>
  <button @click.stop="$router.push({path: '/note', query: {note: note.id, fromFolder: fromFolder}})"
          class="group list-items flex select-none cursor-pointer items-center relative px-6 py-4 pl-4 outline focus:outline-blue-600 rounded-xl">
    <div class="flex items-center gap-2">
      <StickyNote/>
      <div class="line-clamp-1" :title="note.title">
        {{ note.title }}
      </div>
    </div>
    <div class="group-hover:flex absolute hidden gap-2 items-center right-0 pr-2">
      <button @click.stop="renameNote(note.id, note.title)"
              class="p-2 border-2 border-black rounded-lg bg-white hover:bg-black hover:text-white">
        <Pencil size="20"/>
      </button>
      <button @click.stop="deleteNote(note.id)"
              class="p-2 border-2 border-black rounded-lg bg-white hover:bg-black hover:text-white">
        <Trash2 size="20"/>
      </button>
    </div>
  </button>
</template>

<style scoped>

</style>