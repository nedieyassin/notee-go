<script setup lang="ts">
import {DialogWrapper} from "vue3-promise-dialog";
import {onMounted, ref, watch} from "vue";
import {useNotesStore} from "./store/notes";
import {gsap} from "gsap";
import {Flip} from "gsap/Flip";
import {Draggable} from "gsap/Draggable";
import {useRoute, useRouter} from "vue-router";
import FolderContentCard from "./components/cards/FolderContentCard.vue";

gsap.registerPlugin(Flip, Draggable);


const folder = ref(null)
const item = ref('')
const router = useRouter();
const route = useRoute();


onMounted(() => {
  useNotesStore().syncNote()
  gsap.set(folder.value, {yPercent: -1000});
})


const hFolder = () => {
  const id = Number(route.query.folder);
  if (id) {
    item.value = `folder${id}`;
    openFolder(id)
  } else {
    closeFolder();
  }
}

const hcFolder = () => {
  router.push({path: route.path, query: {...route.query, folder: undefined}})
}
watch(() => route.query.folder, (val) => {
  hFolder()
})

onMounted(() => {
  setTimeout(() => hFolder(), 500)
})

const openFolder = (id: number) => {
  // return


  console.log('folder open')
  console.log(document.getElementById(item.value))

  Flip.fit(folder.value, document.getElementById(item.value), {scale: true,});
  const state = Flip.getState(folder.value);
  gsap.set(folder.value, {clearProps: true});
  gsap.set(folder.value, {xPercent: 0, top: "0%", left: "0%", yPercent: 0, visibility: "visible", overflow: "hidden"});
  Flip.from(state, {
    duration: .4,
    ease: "power2.inOut",
    scale: true,
    onComplete: () => {
      gsap.set(folder.value, {overflow: "auto"})
      gsap.to(folder.value, {backgroundColor: "#0008", duration: 0.2})
    } // to permit scrolling if necessary
  })
}

const closeFolder = () => {
  console.log('folder close')
  gsap.set(folder.value, {overflow: "hidden"});
  const state = Flip.getState(folder.value);

  Flip.fit(folder.value, document.getElementById(item.value), {scale: true});

  const tl = gsap.timeline();
  tl.set(folder.value, {overflow: "hidden"})
      .to(folder.value, {backgroundColor: "#0000", duration: 0.2})

  Flip.from(state, {
    scale: true,
    duration: .2,
    opacity: .5,
    delay: 0.1, // 0.2 seconds because we want the details to slide up first, then flip.
    onInterrupt: () => {
      tl.kill()
    },
    onComplete: () => {
      setTimeout(() => hcFolder(), 100)
    }
  }).set(folder.value, {visibility: "hidden"});
}


</script>
<template>
  <DialogWrapper class="z-[999]" :transition-attrs="{name: 'fade'}"/>
  <div ref="folder" class="fixed w-full h-full z-30 p-2 pb-32 md:p-20">
    <FolderContentCard @close="closeFolder"/>
  </div>
  <router-view></router-view>
</template>
