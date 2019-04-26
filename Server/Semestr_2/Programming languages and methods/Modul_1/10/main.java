import java.util.AbstractSet;
import java.util.ArrayList;
import java.util.Collection;
import java.util.Iterator;

/**
 * Created by User on 04.04.2017.
 */
public class SparseSet<T extends Hintable> extends AbstractSet<T> {
    private ArrayList<T> dense;

    public SparseSet() {
        dense = new ArrayList<>();
    }

    @Override
    public boolean contains(Object elem) {
        T x = (T)elem;
        return !dense.isEmpty() && dense.get(x.hint()) == x;
    }

    @Override
    public boolean add(T x) {
        if (!contains(x)) {
            x.setHint(dense.size());
            dense.add(x);
            return true;
        } else {
            return false;
        }
    }

    @Override
    public boolean addAll(Collection<? extends T> col) {
        for (T x: col) {
            if (!add(x)) return false;
        }
        return true;
    }

    @Override
    public boolean remove(Object elem) {
        if (contains(elem)) {
            int n = dense.size();
            T x = (T)elem;
            dense.set(x.hint(), dense.get(n-1));
            dense.get(x.hint()).setHint(x.hint());
            dense.remove(n-1);
            return true;
        } else {
            return false;
        }
    }

    @Override
    public void clear() {
        dense.clear();
    }

    @Override
    public int size() {
        return dense.size();
    }

    @Override
    public Iterator<T> iterator() {
        return new TIterator();
    }

    private class TIterator implements Iterator<T> {
        private int INDEX;

        public TIterator() {
            INDEX = 0;
        }

        @Override
        public boolean hasNext() {
            return INDEX < dense.size();
        }

        @Override
        public T next() {
            return dense.get(INDEX++);
        }

        @Override
        public void remove() {
            SparseSet.this.remove(dense.get(INDEX-1));
        }
    }
}

