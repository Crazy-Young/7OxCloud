/**
 * @description: 展平无限递归树
 * @param {Array} data 无限递归树
 * @param {Object} pData 传递给子元素的父元素数据
 * @returns {Array} 
 * @author: 涂国彬
 */

export default function flattenData(data, pData = {}) {
    const flattenedData = [];
    for (const item of data) {
        const {
            cid,
            children,
            ...rest
        } = item;
        flattenedData.push({
            cid,
            ...pData,
            ...rest
        });

        if (children && children.length > 0) {
            const childrenData = flattenData(children, {
                pid: cid,
                pAuthor: item.author
            });
            flattenedData.push(...childrenData);
        }
    }

    return flattenedData;
}