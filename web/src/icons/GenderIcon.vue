<template>
    <MaleIcon v-if="gender"></MaleIcon>
    <FemaleIcon v-else></FemaleIcon>
</template>

<script>
import MaleIcon from '@/icons/MaleIcon';
import FemaleIcon from '@/icons/FemaleIcon';
export default {
    name: "GenderIcon",
    components: {
        MaleIcon,
        FemaleIcon,
    },
    props: {
        gender: {
            type: Number,
            default: 0
        }
    }
}
</script>