public class Element<T> {
    private Element<T> parent;
    private final T value;

    public Element(T x) {
        this.parent = this;
        this.value = x;
    }

    private Element<T> find(Element<T> element) {
        Element<T> parent = element.parent;
        if (parent.equals(element)) {
            return element;
        }
        Element<T> newParent = find(parent);
        element.parent = newParent;
        return newParent;
    }

    public T x() {
        return this.value;
    }

    public boolean equivalent(Element<T> elem) {
        return find(this) == find(elem);
    }

    public void union(Element<T> elem) {
        Element<T> parent1 = find(this);
        Element<T> parent2 = find(elem);
        if (Math.round(Math.random()) == 1) {
            parent1.parent = parent2;
        } else {
            parent2.parent = parent1;
        }
    }
}
