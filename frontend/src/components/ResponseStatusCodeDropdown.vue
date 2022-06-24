<template>
    <select class="form-control" v-on:input="$emit('updateInput', $event.target.value)">
        <option value="">--- Select ---</option>
        <option v-bind:key="option.Code" v-bind:value="option.Code" v-for="option in options" :selected="this.selected == option.Code">
            [{{ option.Code }}] {{ option.Label }}
        </option>
    </select>
</template>

<script>
export default {
    name: "ResponseStatusCodesDropdown",

    props: {
        options: Array,
        selected: String,
    },
}
</script>