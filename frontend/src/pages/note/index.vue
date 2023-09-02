<script setup lang="ts">
import EditorJS from '@editorjs/editorjs';
import {Save, ChevronLeft} from "lucide-vue-next";
import {onMounted, Ref, ref} from "vue";

// @ts-ignore
import Header from '@editorjs/header';
// @ts-ignore
import Link from '@editorjs/link';
// @ts-ignore
import List from '@editorjs/list';
// @ts-ignore
import Paragraph from '@editorjs/paragraph';
// @ts-ignore
import Quote from '@editorjs/quote';
// @ts-ignore
import Table from '@editorjs/table';
// @ts-ignore
import Warning from '@editorjs/warning';
// @ts-ignore
import Checklist from '@editorjs/checklist';
// @ts-ignore
import Code from '@editorjs/code';


import createGenericInlineTool, {
  UnderlineInlineTool,
} from 'editorjs-inline-tool'
import {onKeyStroke, watchThrottled} from "@vueuse/core";
import {useNotesStore} from "../../store/notes";
import {INote} from "../../types";
import {useRoute, useRouter} from "vue-router";
import {closeDialog} from "vue3-promise-dialog";

const notesStore = useNotesStore();
const route = useRoute();
const router = useRouter();

const note_title = ref('');
const note = ref(undefined) as Ref<INote | undefined>;

watchThrottled(
    note_title,
    () => {
      notesStore.renameNote(note.value!.id, note_title.value);
    },
    {throttle: 500},
)

onMounted(() => {
  const id = Number(route.query.note);
  if (id) {
    notesStore.getNote(id).then((n) => {
      note_title.value = n.title;
      note.value = n;
      initEditor();
    });
  }
})

const initEditor = () => {
  const editor = new EditorJS({
    holder: 'editorjs',
    data: JSON.parse(note.value!.body || '{}'),
    onReady: () => {
      // new Undo({editor});
      // new DragDrop(editor);
    },
    onChange(api, event) {
      editor.save().then((outputData) => {
        notesStore.updateNote(note.value!.id, JSON.stringify(outputData))
      }).catch((error) => {
        console.log('Saving failed: ', error)
      });
    },
    placeholder: 'Write your note...',
    tools: {
      header: {
        class: Header,
        inlineToolbar: true
      },
      link: {
        class: Link,
        inlineToolbar: true
      },
      list: List,
      quote: Quote,
      table: Table,
      warning: Warning,
      checklist: Checklist,
      code: Code,
      paragraph: {
        class: Paragraph,
        inlineToolbar: true,
      },
      bold: {
        class: createGenericInlineTool({
          sanitize: {
            strong: {},
          },
          shortcut: 'CMD+B',
          tagName: 'STRONG',
          toolboxIcon:
              '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M9 12L9 7.1C9 7.04477 9.04477 7 9.1 7H10.4C11.5 7 14 7.1 14 9.5C14 9.5 14 12 11 12M9 12V16.8C9 16.9105 9.08954 17 9.2 17H12.5C14 17 15 16 15 14.5C15 11.7046 11 12 11 12M9 12H11"/></svg>',
        }),
      },
      italic: {
        class: createGenericInlineTool({
          sanitize: {
            i: {},
          },
          shortcut: 'CMD+I',
          tagName: 'I',
          toolboxIcon:
              '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M13.34 10C12.4223 12.7337 11 17 11 17"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M14.21 7H14.2"/></svg>',
        }),
      },
      underline: UnderlineInlineTool,
    },
  })
}
// onKeyStroke('Escape', (e) => {
//   router.back();
// }, {target: document})

</script>

<template>
  <div class="min-h-screen pt-20 md:bg-gray-100-">
    <!--    -->
    <div class="fixed top-0 left-0 w-full bg-white h-16 flex justify-between px-4 py-4 border-b z-50">
      <div class="flex gap-2 items-center">
        <button @click="$router.replace({path:'/', query: {folder: undefined}})"
                class="p-2 rounded-full bg-gray-200 hover:bg-black hover:text-white transition-colors">
          <ChevronLeft/>
        </button>
        <div>
          <input type="text" v-model="note_title" placeholder="Enter note title"
                 class="text-xl font-extrabold ring-0 px-2 py-1 rounded-lg outline-none hover:bg-gray-100 focus:border">
        </div>
      </div>
      <div class="space-x-2 flex items-center gap-2">
        <div>
          <button class="flex gap-2 items-center px-4 py-2 rounded-full bg-black text-white">
            <Save size="20"/>
            <span>Save</span>
          </button>
        </div>
      </div>
    </div>
    <!--    -->
    <div class="md:py-10">
      <div class="bg-white prose py-10 px-6 md:px-10 max-w-screen-lg mx-auto ">
        <div id="editorjs"></div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>