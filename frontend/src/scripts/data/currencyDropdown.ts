const currencyDropdown = (value: string, label: string, flag: string, endpoint: string = '/api/currency.json', staticPath: string = '/static/country/') => {
    return {
        dropdownOpen: false,
        items: [] as { Code: string; Name: string; Flag: string }[],
        filteredItems: [] as { Code: string; Name: string; Flag: string }[],
        searchQuery: '',
        dataLoaded: false,
        errorMessage: '',
        value,
        label,
        flag: `${staticPath}${flag.toLowerCase()}.svg`,

        async loadItems(): Promise<void> {
            try {
                const response = await fetch(endpoint);
                if (!response.ok) throw new Error('Failed to load data.');
                this.items = await response.json();
                this.filteredItems = this.items;
                this.dataLoaded = true;
            } catch (error) {
                console.error(error);
                this.errorMessage = 'Failed to load data.';
            }
        },

        toggleDropdown(): void {
            this.dropdownOpen = !this.dropdownOpen;
            if (this.dropdownOpen && !this.dataLoaded && !this.errorMessage) {
                this.loadItems();
            }
        },

        clearInput(): void {
            this.searchQuery = '';
            this.filteredItems = this.items;
        },

        filterItems(): void {
            const query = this.searchQuery.toLowerCase();
            this.filteredItems = this.items.filter(
                item =>
                    item.Name.toLowerCase().includes(query) ||
                    item.Code.toLowerCase().includes(query)
            );
        },

        selectItem(item: { Code: string; Name: string; Flag: string }): void {
            this.value = item.Code;
            this.label = item.Name;
            this.flag = `${staticPath}${item.Flag.toLowerCase()}.svg`;
            this.dropdownOpen = false;
            window.dispatchEvent(new CustomEvent('currency-selected', { detail: { value: item.Code } }));
        }
    };
}
export default currencyDropdown;