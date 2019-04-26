import java.util.*;

/**
 * Created by User on 06.04.2017.
 */
public class SkipList<K extends Comparable<K>,V> extends AbstractMap<K,V> {
    private int n, levels;
    private Element currrent;

    private class Element implements Entry<K, V> {
        private K key;
        private V value;
        private Element[] next;

        public Element() {
            this.key = null;
            this.value = null;
            this.next = new SkipList.Element[levels];
        }

        public Element(K key, V value) {
            this.key = key;
            this.value = value;
            this.next = new SkipList.Element[levels];
        }

        public K getKey() {
            return key;
        }

        public V getValue() {
            return value;
        }

        public V setValue(V value) {
            return this.value = value;
        }
    }

    public SkipList(int levels) {
        this.levels = levels;
        this.currrent = new Element();
        this.n = 0;
    }

    private Element[] skip(K key) {
        Element x = currrent;
        Element[] p = new SkipList.Element[this.levels];
        for(int i = this.levels-1; i >= 0; i--) {
            while (x.next[i] != null && x.next[i].key.compareTo(key) < 0) {
                x = x.next[i];
            }
            p[i] = x;
        }
        return p;
    }

    public void clear() {
        this.n = 0;
        this.currrent = new Element();
    }

    public V remove(K key) {
        Element[] p = skip(key);
        Element x = p[0].next[0];
        if (x == null || !x.key.equals(key)) {
            return null;
        }
        for(int i = 0; i < this.levels && p[i].next[i] == x; i++) {
            p[i].next[i] = x.next[i];
        }
        n--;
        return x.value;
    }

    public V put(K key, V value) {
        Element[] p = skip(key);
        if (p[0].next[0] != null && p[0].next[0].key.equals(key)) {
            V val = p[0].next[0].value;
            p[0].next[0].setValue(value);
            return val;
        }
        Element x = new Element(key, value);
        Random random = new Random();
        int i, r;
        for (i = 0, r = random.nextInt()*2; i < this.levels && r%2 == 0; i++, r /= 2) {
            x.next[i] = p[i].next[i];
            p[i].next[i] = x;
        }
        for (; i < this.levels; i++) {
            x.next[i] = null;
        }
        n++;
        return null;
    }

    public Set<Entry<K,V>> entrySet() {
        return new SkipListEntrySet();
    }

    private class SkipListEntrySet extends AbstractSet {
        public Iterator<SkipList<K, V>> iterator() {
            return new SkipListIterator();
        }

        private class SkipListIterator implements Iterator {
            private Element i;

            public SkipListIterator() {
                this.i = currrent;
            }

            public boolean hasNext() {
                return i.next[0] != null;
            }

            public Element next() {
                return i = i.next[0];
            }

            public void remove() {
                SkipList.this.remove(i.key);
            }
        }

        public int size() {
            return n;
        }
    }
}
