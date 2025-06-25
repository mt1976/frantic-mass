#Before we can use the script, we have to make it executable with the chmod command:
#chmod +x ./go-executable-build.sh
#then we can use it ./go-executable-build.sh yourpackage
#!/usr/bin/env bash
package="github.com/mt1976/cosmic-orange"
package_name="cosmic-orange"
NOW=$(date +"%y%m")
figlet4go -str "Building"
echo Building Windows x86_64
go-winres simply --icon images/app.png
env GOOS=windows GOARCH=amd64 go build -o "./"$package_name".exe" $package
echo Building Crossplatform

#the full list of the platforms: https://golang.org/doc/install/source#environment
platforms=(
#"darwin/386"
"darwin/amd64"
#"darwin/arm"
"darwin/arm64"
#"dragonfly/amd64"
#"freebsd/386"
#"freebsd/amd64"
#"freebsd/arm"
"linux/386"
"linux/amd64"
"linux/arm"
"linux/arm64"
#"netbsd/386"
#"netbsd/amd64"
#"netbsd/arm"
#"openbsd/386"
#"openbsd/amd64"
#"openbsd/arm"
#"plan9/386"
#"plan9/amd64"
#"solaris/amd64"
"windows/amd64"
"windows/386" )

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    echo Building $GOOS $GOARCH
    go-winres simply --icon images/app.png
    env GOOS=$GOOS GOARCH=$GOARCH go build -o ./exec/$platform/$output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done




echo Done