export default function waterFall(el, columnCount = 5, gap = 20) {
    if (!el) {
        return
    }
    const containerRect = el.getBoundingClientRect()
    const containerWidth = containerRect.width
    const columnWidth = (containerWidth - (columnCount - 1) * gap) / columnCount
    const columnHeight = []
    for (let i = 0; i < columnCount; i++) {
        columnHeight[i] = 0
    }
    const children = el.children
    for (let i = 0; i < children.length; i++) {
        children[i].style.width = columnWidth + 'px'
    }
    for (let i = 0; i < children.length; i++) {
        const child = children[i]
        const childRect = child.getBoundingClientRect()
        const childHeight = childRect.height
        let minHeight = Infinity
        let minIndex = -1
        for (let j = 0; j < columnCount; j++) {
            if (columnHeight[j] < minHeight) {
                minHeight = columnHeight[j]
                minIndex = j
            }
        }
        child.style.left = columnWidth * minIndex + gap * minIndex + 'px'
        child.style.top = columnHeight[minIndex] + 'px'
        columnHeight[minIndex] += childHeight + gap
    }
    el.style.height = Math.max(...columnHeight) + 'px'
}