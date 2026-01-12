<template>
  <div class="contribution-graph">
    <div class="flex flex-col gap-2">
      <!-- Month labels -->
      <div class="flex gap-1 mb-2">
        <div class="w-8"></div>
        <div class="flex-1 flex gap-1">
          <div v-for="month in months" :key="month.index" :style="{ width: month.width + '%' }" class="text-xs text-zinc-400">
            {{ month.name }}
          </div>
        </div>
      </div>

      <div class="flex gap-1">
        <!-- Day labels -->
        <div class="flex flex-col gap-1 text-xs text-zinc-400 justify-around w-8">
          <div>Mon</div>
          <div>Wed</div>
          <div>Fri</div>
        </div>

        <!-- Grid -->
        <div class="flex-1 grid grid-flow-col gap-1" :style="gridStyle">
          <div
            v-for="(day, index) in days"
            :key="index"
            :title="getTooltip(day)"
            :class="getCellClass(day)"
            :style="{ backgroundColor: getCellColor(day) }"
            @click="toggleDay(day)"
            class="aspect-square rounded-sm cursor-pointer hover:ring-2 hover:ring-white/20 transition-all"
          ></div>
        </div>
      </div>

      <!-- Legend -->
      <div class="flex items-center gap-2 justify-end mt-4 text-xs text-zinc-400">
        <span>Less</span>
        <div class="flex gap-1">
          <div class="w-3 h-3 rounded-sm" :style="{ backgroundColor: baseColor + '20' }"></div>
          <div class="w-3 h-3 rounded-sm" :style="{ backgroundColor: baseColor + '60' }"></div>
          <div class="w-3 h-3 rounded-sm" :style="{ backgroundColor: baseColor + 'A0' }"></div>
          <div class="w-3 h-3 rounded-sm" :style="{ backgroundColor: baseColor }"></div>
        </div>
        <span>More</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';

const props = defineProps({
  entries: {
    type: Array,
    default: () => [],
  },
  color: {
    type: String,
    default: '#10b981',
  },
  months: {
    type: Number,
    default: 12,
  },
});

const emit = defineEmits(['toggle']);

const baseColor = computed(() => props.color || '#10b981');

const days = computed(() => {
  const result = [];
  const today = new Date();
  const startDate = new Date(today);
  startDate.setMonth(today.getMonth() - props.months);
  startDate.setDate(1);

  // Start from the first Sunday before or on the start date
  const dayOfWeek = startDate.getDay();
  startDate.setDate(startDate.getDate() - dayOfWeek);

  const currentDate = new Date(startDate);
  while (currentDate <= today) {
    const dateStr = currentDate.toISOString().split('T')[0];
    const entry = props.entries.find(e => e.date.startsWith(dateStr));
    
    result.push({
      date: new Date(currentDate),
      dateStr,
      completed: entry?.completed || false,
      dayOfWeek: currentDate.getDay(),
    });

    currentDate.setDate(currentDate.getDate() + 1);
  }

  return result;
});

const months = computed(() => {
  const monthNames = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
  const result = [];
  const weeks = Math.ceil(days.value.length / 7);
  
  let currentMonth = -1;
  let monthStartWeek = 0;
  
  for (let i = 0; i < days.value.length; i += 7) {
    const weekMonth = days.value[i].date.getMonth();
    if (weekMonth !== currentMonth) {
      if (currentMonth !== -1) {
        const weeksPassed = i / 7 - monthStartWeek;
        result.push({
          name: monthNames[currentMonth],
          width: (weeksPassed / weeks) * 100,
          index: currentMonth,
        });
      }
      currentMonth = weekMonth;
      monthStartWeek = i / 7;
    }
  }
  
  // Add the last month
  if (currentMonth !== -1) {
    const weeksPassed = (days.value.length / 7) - monthStartWeek;
    result.push({
      name: monthNames[currentMonth],
      width: (weeksPassed / weeks) * 100,
      index: currentMonth,
    });
  }
  
  return result;
});

const gridStyle = computed(() => {
  const cols = Math.ceil(days.value.length / 7);
  return {
    gridTemplateRows: 'repeat(7, 1fr)',
    gridTemplateColumns: `repeat(${cols}, 1fr)`,
  };
});

function getCellClass(day) {
  const baseClass = 'border border-zinc-800';
  return day.completed ? baseClass : baseClass;
}

function getCellColor(day) {
  if (!day.completed) {
    return 'rgba(39, 39, 42, 0.5)'; // zinc-800 with opacity
  }
  return baseColor.value;
}

function getTooltip(day) {
  const dateStr = day.date.toLocaleDateString('en-US', { 
    weekday: 'short', 
    year: 'numeric', 
    month: 'short', 
    day: 'numeric' 
  });
  return day.completed ? `Completed on ${dateStr}` : `No activity on ${dateStr}`;
}

function toggleDay(day) {
  emit('toggle', day.dateStr);
}
</script>

<style scoped>
.contribution-graph {
  @apply w-full;
}
</style>
