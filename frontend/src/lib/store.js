import { writable } from 'svelte/store';

export const isAuthenticated = writable(false);
export const contentLoading = writable(false);
export const settings = writable({
    pomodoro: 25,
    shortBreak: 5,
    longBreak: 30,
    theme: "Default",
    sound: "Bell",
    volume: "High",
});
export const pomodoros = writable([]);
