<script setup lang="ts">
import {onMounted, Ref, ref, watch} from "vue";
import {INote} from "../../types";
import {Folders, StickyNote, Plus, CircleOff} from "lucide-vue-next";
import {onClickOutside, watchThrottled} from "@vueuse/core";
import {useNotesStore} from "../../store/notes";
import {useRoute, useRouter} from "vue-router";
import FolderCard from "./FolderCard.vue";
import NoteCard from "./NoteCard.vue";

const notesStore = useNotesStore();
const router = useRouter();
const route = useRoute();

const note_title = ref('');
const pnote = ref(undefined) as Ref<INote | undefined>;
const notes = ref([]) as Ref<INote[]>;

const target = ref(null)


const emit = defineEmits(['close'])

watchThrottled(
    note_title,
    () => {
      if (route.query.folder) {
        notesStore.renameNote(pnote.value!.id, note_title.value);
      }
    },
    {throttle: 500},
)

onClickOutside(target, (event) => {
  if (route.query.folder) {
    emit('close')
  }
})

watch(() => route.query, () => {
  const id = Number(route.query.folder);
  if (id) {
    notesStore.getNote(id).then((n) => {
      note_title.value = n.title;
      pnote.value = n;
      notes.value = notesStore.getFolderNotes(id);
    });
  } else {
    note_title.value = '';
    pnote.value = undefined
  }
})


const createNote = () => {
  notesStore.createNote(pnote.value!.id).then((id) => {
    if (id) {
      router.push({path: '/note', query: {note: id}})
    }
  });
}

</script>

<template>
  <div ref="target"
       class="bg-white w-full h-full relative md:mx-0 z-50 rounded-xl shadow-xl">
    <div class="absolute top-0 left-0 w-full h-16 flex justify-between px-4 py-4 z-50">
      <div class="flex gap-2 items-center">
        <Folders class="text-blue-600"/>
        <div>
          <input type="text" v-model="note_title" placeholder="Enter note title"
                 class="text-xl font-extrabold ring-0 px-2 py-1 rounded-lg outline-none hover:bg-gray-100 focus:border">
        </div>
      </div>
      <div class="space-x-2 flex items-center gap-2">
        <div>
          <button @click="createNote"
                  class="hover:opacity-75 transition-all flex gap-2 items-center px-4 py-2 rounded-full bg-black text-white">
            <Plus size="20" class="block md:hidden"/>
            <StickyNote size="20" class="hidden md:block"/>
            <span class="hidden md:block">Add Note</span>
          </button>
        </div>
      </div>
    </div>

    <div class="pt-20">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-6 gap-3 px-6 py-4">
        <template v-for="note in notes">
          <NoteCard
              :from-folder="pnote?.id"
              tabindex="0"
              :note="note"/>
        </template>
      </div>
      <div v-if="!notes?.length"
           class="select-none cursor-default flex py-16 text-gray-300 flex-col w-full justify-center items-center gap-4">
        <CircleOff size="100"/>
        <h1 class="text-2xl">No notes found</h1>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>