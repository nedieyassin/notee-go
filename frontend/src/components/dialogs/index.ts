import {openDialog} from 'vue3-promise-dialog'
import PromptDialog from "./PromptDialog.vue";
import ConfirmDialog from "./ConfirmDialog.vue";


export const promptDialog = (args: Record<any, any>) => openDialog(PromptDialog as any, {...args});
export const confirmDialog = (args: Record<any, any>) => openDialog(ConfirmDialog as any, {...args});