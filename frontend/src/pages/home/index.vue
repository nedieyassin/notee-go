<script setup lang="ts">
import {StickyNote, CircleOff} from 'lucide-vue-next';
import NoteCard from "../../components/cards/NoteCard.vue";
import {useNotesStore} from "../../store/notes";
import FolderCard from "../../components/cards/FolderCard.vue";
import {promptDialog} from "../../components/dialogs";
import {useRoute, useRouter} from "vue-router";
import FolderContentCard from "../../components/cards/FolderContentCard.vue";
//

//
import {onMounted, ref, watch} from "vue";
import {onKeyStroke} from "@vueuse/core";


const notesStore = useNotesStore();
const router = useRouter();
const route = useRoute();


const createFolder = () => {
  promptDialog({
    title: 'Create new folder',
    placeholder: 'Enter folder name',
    value: 'New folder',
    action: 'Create'
  }).then((res) => {
    if (res) {
      notesStore.createFolder(res);
    }
  })
}

const createNote = () => {
  notesStore.createNote(0).then((id) => {
    if (id) {
      router.push({path: '/note', query: {note: id}})
    }
  });
}

const handleFolder = (id: number) => {
  router.push({path: route.path, query: {folder: id}})
}


</script> :

<template>
  {{$route.query}}
  <div>
    <div class="flex justify-between px-6 py-4">
      <div>
        <h1 class="text-2xl font-extrabold cursor-default select-none">Notee</h1>
      </div>
      <div class="space-x-2 flex items-center gap-2">
        <div>
          <button @click="createFolder"
                  tabindex="0"
                  class="hover:opacity-75 transition-all flex gap-1 items-center px-4 py-2 bg-gray-200 rounded-full text-black">
            <Folder size="20"/>
            <span>Add Folder</span>
          </button>
        </div>
        <div>
          <button @click="createNote"
                  tabindex="0"
                  class="hover:opacity-75 transition-all flex gap-2 items-center px-4 py-2 rounded-full bg-black text-white">
            <StickyNote size="20"/>
            <span>Add Note</span>
          </button>
        </div>
      </div>
    </div>
    <div>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-6 gap-3 px-6 py-4">
        <template v-for="note in notesStore.nodes.filter((n) => n.parentId === 0)">
          <FolderCard
              :id="`folder${note.id}`"
              @click.stop="handleFolder(note.id)"
              :note="note" v-if="note.isDir"/>
          <NoteCard
              :id="`note${note.id}`"
              :note="note" v-else/>
        </template>
      </div>
      <div v-if="!notesStore.nodes?.length"
           class="select-none cursor-default flex py-16 text-gray-300 flex-col w-full justify-center items-center gap-4">
        <CircleOff size="100"/>
        <h1 class="text-2xl">No notes found</h1>
      </div>
    </div>
  </div>


</template>

<style scoped>

</style>