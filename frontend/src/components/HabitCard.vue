<template>
  <div class="habit-card bg-zinc-900 rounded-lg border border-zinc-800 p-6 hover:border-zinc-700 transition-colors">
    <div class="flex items-start justify-between mb-4">
      <div class="flex-1">
        <div class="flex items-center gap-3">
          <div class="w-3 h-3 rounded-full" :style="{ backgroundColor: habit.color }"></div>
          <h3 class="text-lg font-semibold text-white">{{ habit.name }}</h3>
        </div>
        <p v-if="habit.description" class="text-sm text-zinc-400 mt-2 ml-6">
          {{ habit.description }}
        </p>
      </div>
      <div class="flex gap-2">
        <button
          @click="$emit('edit')"
          class="p-2 text-zinc-400 hover:text-white transition-colors"
          title="Edit habit"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
          </svg>
        </button>
        <button
          @click="$emit('delete')"
          class="p-2 text-zinc-400 hover:text-red-400 transition-colors"
          title="Delete habit"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
        </button>
      </div>
    </div>

    <ContributionGraph
      :entries="habit.entries || []"
      :color="habit.color"
      :months="12"
      @toggle="handleToggle"
    />

    <div class="mt-4 flex items-center justify-between text-sm">
      <div class="text-zinc-400">
        <span class="text-white font-medium">{{ completedCount }}</span> days completed
      </div>
      <div class="text-zinc-400">
        Current streak: <span class="text-white font-medium">{{ currentStreak }}</span> days
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import ContributionGraph from './ContributionGraph.vue';

const props = defineProps({
  habit: {
    type: Object,
    required: true,
  },
});

const emit = defineEmits(['toggle', 'edit', 'delete']);

const completedCount = computed(() => {
  return props.habit.entries?.filter(e => e.completed).length || 0;
});

const currentStreak = computed(() => {
  if (!props.habit.entries || props.habit.entries.length === 0) return 0;
  
  const sortedEntries = [...props.habit.entries]
    .filter(e => e.completed)
    .sort((a, b) => new Date(b.date) - new Date(a.date));
  
  if (sortedEntries.length === 0) return 0;
  
  let streak = 0;
  const today = new Date();
  today.setHours(0, 0, 0, 0);
  
  for (let i = 0; i < sortedEntries.length; i++) {
    const entryDate = new Date(sortedEntries[i].date);
    entryDate.setHours(0, 0, 0, 0);
    
    const expectedDate = new Date(today);
    expectedDate.setDate(today.getDate() - i);
    expectedDate.setHours(0, 0, 0, 0);
    
    if (entryDate.getTime() === expectedDate.getTime()) {
      streak++;
    } else {
      break;
    }
  }
  
  return streak;
});

function handleToggle(date) {
  emit('toggle', date);
}
</script>
