<template>
  <div v-if="isOpen" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4" @click.self="close">
    <div class="bg-zinc-900 rounded-lg border border-zinc-800 p-6 w-full max-w-md">
      <h2 class="text-xl font-semibold text-white mb-4">
        {{ habit ? 'Edit Habit' : 'New Habit' }}
      </h2>

      <form @submit.prevent="handleSubmit">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-zinc-300 mb-2">
              Habit Name
            </label>
            <input
              v-model="formData.name"
              type="text"
              required
              class="w-full px-3 py-2 bg-zinc-800 border border-zinc-700 rounded-md text-white placeholder-zinc-500 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
              placeholder="e.g., Morning Exercise"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-zinc-300 mb-2">
              Description (Optional)
            </label>
            <textarea
              v-model="formData.description"
              rows="3"
              class="w-full px-3 py-2 bg-zinc-800 border border-zinc-700 rounded-md text-white placeholder-zinc-500 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent resize-none"
              placeholder="Add a description..."
            ></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-zinc-300 mb-2">
              Color
            </label>
            <div class="flex gap-2 flex-wrap">
              <button
                v-for="color in colors"
                :key="color"
                type="button"
                @click="formData.color = color"
                :class="[
                  'w-8 h-8 rounded-full border-2 transition-all',
                  formData.color === color ? 'border-white scale-110' : 'border-transparent'
                ]"
                :style="{ backgroundColor: color }"
              ></button>
            </div>
          </div>
        </div>

        <div class="flex gap-3 mt-6">
          <button
            type="button"
            @click="close"
            class="flex-1 px-4 py-2 bg-zinc-800 text-white rounded-md hover:bg-zinc-700 transition-colors"
          >
            Cancel
          </button>
          <button
            type="submit"
            class="flex-1 px-4 py-2 bg-emerald-600 text-white rounded-md hover:bg-emerald-700 transition-colors"
          >
            {{ habit ? 'Update' : 'Create' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  isOpen: {
    type: Boolean,
    default: false,
  },
  habit: {
    type: Object,
    default: null,
  },
});

const emit = defineEmits(['close', 'submit']);

const colors = [
  '#10b981', // emerald
  '#3b82f6', // blue
  '#8b5cf6', // violet
  '#ec4899', // pink
  '#f59e0b', // amber
  '#ef4444', // red
  '#06b6d4', // cyan
  '#84cc16', // lime
];

const formData = ref({
  name: '',
  description: '',
  color: '#10b981',
});

watch(() => props.habit, (newHabit) => {
  if (newHabit) {
    formData.value = {
      name: newHabit.name || '',
      description: newHabit.description || '',
      color: newHabit.color || '#10b981',
    };
  } else {
    formData.value = {
      name: '',
      description: '',
      color: '#10b981',
    };
  }
}, { immediate: true });

function handleSubmit() {
  emit('submit', { ...formData.value });
}

function close() {
  emit('close');
}
</script>
