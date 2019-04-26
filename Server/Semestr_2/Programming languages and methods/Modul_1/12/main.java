import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.util.*;

/**
 * Created by User on 08.04.2017.
 */
public class Sync {
    private ArrayList<String> Copy, Delete;
    private final String pathD, pathS;

    public Sync(String pathD, String pathS) {
        this.pathD = pathD;
        this.pathS = pathS;
        this.Copy = new ArrayList<>();
        this.Delete = new ArrayList<>();
    }

    public static void main(String[] args) throws IOException {
        String pathS = args[0], pathD = args[1];
        Sync sync = new Sync(pathD, pathS);
        sync.synchronization("");
        Collections.sort(sync.Delete);
        Collections.sort(sync.Copy);
        for (String str : sync.Delete) {
            System.out.println("DELETE " + str);
        }
        for (String str : sync.Copy) {
            System.out.println("COPY " + str);
        }
        if (sync.Copy.size() == 0 && sync.Delete.size() == 0) {
            System.out.println("IDENTICAL");
        }
    }

    private void synchronization(String path) throws IOException {
        File fileD, fileS;
        if (path == "") {
            fileD = new File(pathD);
            fileS = new File(pathS);
        } else {
            fileD = new File(pathD + '/' + path);
            fileS = new File(pathS + '/' + path);
            path += '/';
        }
        List<String> dateD = Arrays.asList(fileD.list());
        List<String> dateS = Arrays.asList(fileS.list());

        for (String str : dateD) {
            if (dateS.contains(str)) {
                File fileInD = new File(fileD.getCanonicalPath()+'/'+str);
                File fileInS = new File(fileS.getCanonicalPath()+'/'+str);
                if (fileInD.isFile() && fileInS.isFile()) {
                    if (!equalFile(fileInD, fileInS)) {
                        this.Delete.add(path+str);
                        this.Copy.add(path+str);
                    }
                    continue;
                }
                if (fileInD.isDirectory() && fileInS.isDirectory()) {
                    synchronization(path+str);
                    continue;
                }
                if (fileInD.isFile() && fileInS.isDirectory()) {
                    //Удаление файла из D
                    this.Delete.add(path+str);
                    //Копирование содержимого папки из S
                    actionDir(path+str, true);
                    continue;
                }
                if (fileInD.isDirectory() && fileInS.isFile()) {
                    //Удаление папки в D
                    actionDir(path+str, false);
                    //Копирование файла из S
                    this.Copy.add(path+str);
                    continue;
                }
                error();
            } else {
                File fileInD = new File(fileD.getCanonicalPath()+'/'+str);
                if (fileInD.isFile()) {
                    this.Delete.add(path+str);
                    continue;
                }
                if (fileInD.isDirectory()) {
                    actionDir(path+str, false);
                    continue;
                }
                error();
            }
        }

        for (String str : dateS) {
            if (!dateD.contains(str)) {
                File fileInS = new File(fileS.getCanonicalPath()+'/'+str);
                if (fileInS.isFile()) {
                    this.Copy.add(path+str);
                    continue;
                }
                if (fileInS.isDirectory()) {
                    actionDir(path+str, true);
                    continue;
                }
                error();
            }
        }
    }

    private void actionDir(String path, boolean com) throws IOException {
        //com = true => COPY
        //com = false => DELETE
        ArrayList<String> arr;
        if (com) {
            arr = this.Copy;
        } else {
            arr = this.Delete;
        }
        File fileS = new File(pathS+'/'+path);
        List<String> dateS = Arrays.asList(fileS.list());
        for (String str : dateS) {
            File fileInS = new File(fileS.getCanonicalPath()+'/'+str);
            if (fileInS.isFile()) {
                arr.add(path+'/'+str);
                continue;
            }
            if (fileInS.isDirectory()) {
                actionDir(path+'/'+str, com);
            }
        }
    }

    private boolean equalFile(File fileD, File fileS) throws IOException {
        return fileD.length() == fileS.length() && Arrays.equals(Files.readAllBytes(fileD.toPath()), Files.readAllBytes(fileS.toPath()));
    }
    
    private void error() {
        System.exit(1);
    }
}