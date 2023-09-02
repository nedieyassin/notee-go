<script lang="ts" setup>
import {closeDialog} from 'vue3-promise-dialog'
import {onKeyStroke} from "@vueuse/core";

defineProps({
  title: {
    type: String,
    default: 'Prompt',
  },
  message: {
    type: String,
    default: 'Prompt',
  },
  action: {
    type: String,
    default: 'Done',
  },
});

onKeyStroke('Escape', (e) => {
  closeDialog(false)
}, {target: document})

onKeyStroke('Enter', (e) => {
  closeDialog(true)
}, {target: document})

</script>

<template>
  <Transition name="fade">
    <div class="flex fixed justify-center top-0 items-center w-screen h-screen">
      <div @click="closeDialog(false)"
           class="fixed top-0  bg-black/50 left-0 h-screen w-screen  z-50">
      </div>
      <div
          class="mb-96 border   border-gray-200 bg-white w-96 relative mx-3 md:mx-0 z-50 rounded-xl shadow-xl p-2">
        <div class="flex justify-between items-center pl-2">
          <div>
            <h1 class="font-bold text-gray-900">{{ title }}</h1>
          </div>
          <button>
            <svg @click="closeDialog(false)" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"
                 width="24"
                 height="24"
                 class="h-10 w-10  hover:bg-gray-200 text-gray-800 hover:text-red-700  transition-all rounded-full p-1">
              <path
                  fill="currentColor"
                  d="M12 10.586l4.95-4.95 1.414 1.414-4.95 4.95 4.95 4.95-1.414 1.414-4.95-4.95-4.95 4.95-1.414-1.414 4.95-4.95-4.95-4.95L7.05 5.636z"/>
            </svg>
          </button>
        </div>
        <div class="px-2 py-2">
          <p>{{ message }}</p>
        </div>
        <div class="flex gap-3 justify-end pt-1">
          <div>
            <button @click="closeDialog(false)"
                    class="px-4 py-1 bg-gray-200 rounded-lg hover:opacity-75 transition-all">
              Cancel
            </button>
          </div>
          <div>
            <button @click="closeDialog(true)"
                    class="px-4 py-1 bg-black text-white rounded-lg hover:opacity-75 transition-all">
              {{ action }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>