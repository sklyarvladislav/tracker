const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

export const habitAPI = {
  // Habits
  async getHabits() {
    const response = await fetch(`${API_BASE_URL}/habits`);
    if (!response.ok) throw new Error('Failed to fetch habits');
    return response.json();
  },

  async getHabit(id) {
    const response = await fetch(`${API_BASE_URL}/habits/${id}`);
    if (!response.ok) throw new Error('Failed to fetch habit');
    return response.json();
  },

  async createHabit(habit) {
    const response = await fetch(`${API_BASE_URL}/habits`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(habit),
    });
    if (!response.ok) throw new Error('Failed to create habit');
    return response.json();
  },

  async updateHabit(id, habit) {
    const response = await fetch(`${API_BASE_URL}/habits/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(habit),
    });
    if (!response.ok) throw new Error('Failed to update habit');
    return response.json();
  },

  async deleteHabit(id) {
    const response = await fetch(`${API_BASE_URL}/habits/${id}`, {
      method: 'DELETE',
    });
    if (!response.ok) throw new Error('Failed to delete habit');
    return response.json();
  },

  // Entries
  async getHabitEntries(habitId) {
    const response = await fetch(`${API_BASE_URL}/habits/${habitId}/entries`);
    if (!response.ok) throw new Error('Failed to fetch entries');
    return response.json();
  },

  async toggleHabitEntry(habitId, date) {
    const response = await fetch(`${API_BASE_URL}/habits/${habitId}/toggle?date=${date}`, {
      method: 'POST',
    });
    if (!response.ok) throw new Error('Failed to toggle entry');
    return response.json();
  },

  async createHabitEntry(habitId, entry) {
    const response = await fetch(`${API_BASE_URL}/habits/${habitId}/entries`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(entry),
    });
    if (!response.ok) throw new Error('Failed to create entry');
    return response.json();
  },
};
