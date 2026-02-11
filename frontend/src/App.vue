<script setup>
import { ref, onMounted, nextTick } from 'vue';
import HabitCard from './components/HabitCard.vue';
import HabitForm from './components/HabitForm.vue';
import { habitAPI } from './services/api.js';

const habits = ref([]);
const loading = ref(true);
const error = ref(null);
const isFormOpen = ref(false);
const editingHabit = ref(null);

onMounted(async () => {
  await loadHabits();
});

async function loadHabits() {
  try {
    loading.value = true;
    error.value = null;
    habits.value = await habitAPI.getHabits();
  } catch (err) {
    error.value = err.message;
    console.error('Failed to load habits:', err);
  } finally {
    loading.value = false;
  }
}

function openNewHabitForm() {
  editingHabit.value = null;
  isFormOpen.value = true;
}

function openEditHabitForm(habit) {
  editingHabit.value = habit;
  isFormOpen.value = true;
}

function closeForm() {
  isFormOpen.value = false;
  editingHabit.value = null;
}

async function handleFormSubmit(habitData) {
  try {
    if (editingHabit.value) {
      await habitAPI.updateHabit(editingHabit.value.id, habitData);
    } else {
      await habitAPI.createHabit(habitData);
    }
    closeForm();
    await loadHabits();
  } catch (err) {
    error.value = err.message;
    console.error('Failed to save habit:', err);
  }
}

async function handleToggleEntry(habitId, date) {
  try {
    // Save scroll position before reloading
    const scrollPosition = window.scrollY;
    
    await habitAPI.toggleHabitEntry(habitId, date);
    await loadHabits();
    
    // Restore scroll position after DOM updates
    await nextTick();
    window.scrollTo(0, scrollPosition);
  } catch (err) {
    error.value = err.message;
    console.error('Failed to toggle entry:', err);
  }
}

async function handleDeleteHabit(habitId) {
  if (!confirm('Are you sure you want to delete this habit?')) {
    return;
  }

  try {
    await habitAPI.deleteHabit(habitId);
    await loadHabits();
  } catch (err) {
    error.value = err.message;
    console.error('Failed to delete habit:', err);
  }
}
</script>

<template>
  <div class="min-h-screen bg-zinc-950 dark">
    <!-- Header -->
    <header class="border-b border-zinc-800 bg-zinc-900">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-white">Tracker</h1>
            <p class="text-zinc-400 mt-1">Track your daily habits with contribution graph</p>
          </div>
          <button
            @click="openNewHabitForm"
            class="px-4 py-2 bg-emerald-600 text-white rounded-md hover:bg-emerald-700 transition-colors flex items-center gap-2"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            New Habit
          </button>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Error Message -->
      <div v-if="error" class="mb-6 p-4 bg-red-900/20 border border-red-900 rounded-lg text-red-400">
        {{ error }}
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex items-center justify-center py-12">
        <div class="text-zinc-400">Loading habits...</div>
      </div>

      <!-- Empty State -->
      <div v-else-if="habits.length === 0" class="text-center py-12">
        <svg class="w-16 h-16 mx-auto text-zinc-700 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
        </svg>
        <h2 class="text-xl font-semibold text-white mb-2">No habits yet</h2>
        <p class="text-zinc-400 mb-4">Start tracking your daily habits by creating your first one</p>
        <button
          @click="openNewHabitForm"
          class="px-4 py-2 bg-emerald-600 text-white rounded-md hover:bg-emerald-700 transition-colors"
        >
          Create Your First Habit
        </button>
      </div>

      <!-- Habits List -->
      <div v-else class="space-y-6">
        <HabitCard
          v-for="habit in habits"
          :key="habit.id"
          :habit="habit"
          @toggle="(date) => handleToggleEntry(habit.id, date)"
          @edit="openEditHabitForm(habit)"
          @delete="handleDeleteHabit(habit.id)"
        />
      </div>
    </main>

    <!-- Habit Form Modal -->
    <HabitForm
      :is-open="isFormOpen"
      :habit="editingHabit"
      @close="closeForm"
      @submit="handleFormSubmit"
    />
  </div>
</template>

<style scoped>
</style>

