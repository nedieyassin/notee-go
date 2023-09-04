<script setup lang="ts">
import {Folders, MoreVertical, Pencil, Star, Trash2, X} from "lucide-vue-next";
import {onMounted, PropType, ref} from "vue";
import {INote} from "../../types";
import {useNotesStore} from "../../store/notes";
import {confirmDialog, promptDialog} from "../dialogs";


defineProps({
  note: {
    type: Object as PropType<INote>,
    required: true
  }
})

const notesStore = useNotesStore();
const isOpen = ref(false);


const deleteNote = (id: number) => {
  confirmDialog({
    title: 'Delete Folder?',
    message: "Are you sure you want to delete this folder and it's notes?",
    action: 'Yes, Delete'
  }).then((res) => {
    if (res) {
      notesStore.deleteNote(id)
    }
  })
}

const renameNote = (id: number, title: string) => {
  promptDialog({
    title: 'Edit folder name',
    placeholder: 'Enter note name',
    action: 'Edit',
    value: title
  }).then((res) => {
    if (res) {
      notesStore.renameNote(id, res);
    }
  })
}

const isFav = (id: number, isFav: number) => {
  notesStore.updateFav(id, isFav === 1 ? 0 : 1);
}

</script>

<template>
  <button ref="item"
          class="group list-items flex select-none cursor-pointer items-center relative px-6 py-4 pl-4 outline focus:outline-blue-600  rounded-xl">
    <div class="flex items-center gap-2">
      <Folders class="text-blue-600"/>
      <div class="line-clamp-1">
        {{ note?.title }}
      </div>
    </div>
    <div class="group-hover:flex absolute hidden gap-2 items-center right-0 pr-2">
      <button v-if="!isOpen" @click.stop="isFav(note.id, note.isFav)"
              :class="[note.isFav ? 'bg-black text-white' : 'bg-white hover:bg-black hover:text-white']"
              class="p-2 border-2 border-black rounded-lg">
        <Star size="20"/>
      </button>
      <button v-if="isOpen" @click.stop="renameNote(note.id, note.title)"
              class="p-2 border-2 border-black rounded-lg bg-white hover:bg-black hover:text-white">
        <Pencil size="20"/>
      </button>
      <button v-if="isOpen" @click.stop="deleteNote(note.id)"
              class="p-2 border-2 border-black rounded-lg bg-white hover:bg-black hover:text-white">
        <Trash2 size="20"/>
      </button>
      <button @click.stop="isOpen = !isOpen"
              class="p-2 border-2 border-black rounded-lg bg-white hover:bg-black hover:text-white">
        <X v-if="isOpen" size="20"/>
        <MoreVertical v-else size="20"/>
      </button>
    </div>
  </button>
</template>

<style scoped>

</style>