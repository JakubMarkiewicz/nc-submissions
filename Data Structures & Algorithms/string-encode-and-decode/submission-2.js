class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {
        return strs.map(v => `__${v}__`).join("-||-")
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        return str.split("-||-").filter(v => v.startsWith("__") && v.endsWith("__")).map(v => v.slice(2,-2))
    }
}
