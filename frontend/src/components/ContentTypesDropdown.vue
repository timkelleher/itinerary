<template>
    <select class="form-control" v-on:input="$emit('updateInput', $event.target.value)">
        <option v-bind:key="option" v-for="option in options" :selected="this.selected == option">
            {{ option }}
        </option>
    </select>
</template>

<script>
export default {
    name: "ContentTypesDropdown",

    props: {
        options: Array,
        selected: String,
    },
}
</script>