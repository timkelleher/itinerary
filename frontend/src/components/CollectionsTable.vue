<template>
    <div class="col-md-10 offset-md-1 py-2">
        <div class="row">
            <div class="col-8">
                <h3>Collections</h3>
            </div>
            <div class="col-4 text-right">
                <router-link :to="{ name: 'collection_add' }">
                    <button class="btn btn-primary btn-sm">New Collection</button>
                </router-link>
            </div>
        </div>
        <div class="row">
        <table class="table table-bordered">
            <tr>
                <th>Name</th>
                <th>Endpoints</th>
                <th>Path Prefix</th>
            </tr>
            <tr v-bind:key="collection.ID" v-for="collection in collections">
                <td>
                    <router-link :to="{ name: 'collection_detail', params: { id: collection.ID } }" >
                        {{ collection.Name }}
                    </router-link>
                </td>
                <td>17</td>
                <td>{{ collection.PathPrefix }}</td>
            </tr>
        </table>
        </div>
        <router-view />
    </div>
</template>

<script>
export default {
    name: "CollectionsTable",

    props: {
        "collections": Array,
    },
}
</script>