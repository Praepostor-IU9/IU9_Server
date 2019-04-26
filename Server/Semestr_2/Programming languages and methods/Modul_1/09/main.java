
import java.util.AbstractSet;
import java.util.Iterator;

/**
 * Created by User on 03.04.2017.
 */
public class IntSparseSet extends AbstractSet<Integer> {
    private int u, n, low;
    private int[] sparse, dense;

    public IntSparseSet(int low, int high) {
        this.low = low;
        u = high-low;
        n = 0;
        sparse = new int[u];
        dense = new int[u];
    }

    @Override
    public boolean contains(Object elem) {
        int x = (int) elem;
        return x >= low && x < u + low && sparse[x - low] < n && dense[sparse[x - low]] == x;
    }

    @Override
    public boolean add(Integer x) {
        if (x >= low && x < u + low && !(sparse[x - low] < n && dense[sparse[x - low]] == x)) {
            sparse[x - low] = n;
            dense[n] = x;
            n++;
            return true;
        } else {
            return false;
        }
    }

    @Override
    public boolean remove(Object elem) {
        int x = (int) elem;
        if (contains(elem)) {
            dense[sparse[x - low]] = dense[n - 1];
            sparse[dense[n - 1]-low] = sparse[x - low];
            n--;
            return true;
        } else {
            return false;
        }
    }

    @Override
    public void clear() {
        n = 0;
    }

    @Override
    public int size() {
        return n;
    }

    @Override
    public Iterator<Integer> iterator() {
        return new IntIterator();
    }

    private class IntIterator implements Iterator<Integer> {
        private int INDEX;

        public IntIterator() {
            INDEX = 0;
        }

        @Override
        public boolean hasNext() {
            return INDEX < n;
        }

        @Override
        public Integer next() {
            return dense[INDEX++];
        }
        
        @Override
        public void remove() {
            dense[INDEX - 1] = dense[n-1];
            sparse[dense[n-1]-low] = INDEX-1;
            n--;
        }
    }
}