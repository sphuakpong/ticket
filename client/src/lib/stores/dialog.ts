import { type ComponentType } from 'svelte';
import { writable, type Writable } from 'svelte/store';

type Data = any;

export type DialogState<T> = {
	initializing: boolean;
	open: boolean;
	forceClose: boolean;
	component: ComponentType | null;
	data?: T;
	onOpenChange: (isOpen: boolean) => void;
	onClose?: (d: T) => boolean;
};

function onOpenChange(isOpen: boolean) {
	console.log('onOpenChange', isOpen);
	dialogStore.update((store) => {
		if (!isOpen && !store.forceClose && store.onClose) {
			const closeable = store.onClose(store.data);
			store.open = closeable == false;
		} else {
			store.open = isOpen;
		}

		return store;
	});
}

function defaultState(): DialogState<Data> {
	return {
		initializing: true,
		open: false,
		component: null,
		forceClose: false,
		onOpenChange: onOpenChange
	};
}

const dialogStore = writable<DialogState<any>>(defaultState());

export interface DialogStore<T> extends Writable<DialogState<T>> {
	state: DialogState<T>;
}

function create(comp: ComponentType) {
	dialogStore.update((store) => {
		if (comp !== store.component) {
			console.log('create', comp);
			store.component = comp;
		}

		store.open = true;

		return store;
	});
}

function close() {
	dialogStore.update((store) => {
		store.forceClose = true;
		store.open = false;
		setTimeout(() => {
			dialogStore.update((store) => {
				store.component = null;

				return store;
			});
		}, 100);

		return store;
	});
}

export const DialogStore = {
	...dialogStore,
	defaultState,
	create,
	close
};
